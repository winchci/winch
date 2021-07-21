/*
winch - Universal Build and Release Tool
Copyright (C) 2021 Ketch Kloud, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package pkg

import (
	"context"
	"time"
)

type ReleaseSection struct {
	Title      string
	IsBreaking bool
	Commits    []*Commit
}

type ChangelogRelease struct {
	Version         string
	PreviousVersion string
	Date            string
	Timestamp       time.Time
	Sections        []*ReleaseSection
	FirstCommitHash string
	LastCommitHash  string
	IsNew           bool
}

type Changelog struct {
	Repository string
	Releases   []*ChangelogRelease
}

func MakeChangelog(_ context.Context, repository string, releases []*Release) (*Changelog, error) {
	c := &Changelog{
		Repository: repository,
	}

	for _, r := range releases {
		clr := &ChangelogRelease{
			Version:         r.Version,
			PreviousVersion: r.PreviousVersion,
			Date:            r.Date.Format("2006-01-02"),
			Timestamp:       r.Date,
			Sections:        nil,
			FirstCommitHash: r.FirstCommitHash,
			LastCommitHash:  r.LastCommitHash,
			IsNew:           r.IsNew,
		}

		sections := make(map[string][]*Commit)
		for _, vv := range r.Commits {
			for _, v := range vv {
				if v.Message.IsBreaking {
					sections[breakingChangeTitle] = append(sections[breakingChangeTitle], v)
				}

				title := v.Title()
				sections[title] = append(sections[title], v)
			}
		}

		for _, k := range titleOrder {
			if v, ok := sections[k]; ok {
				clr.Sections = append(clr.Sections, &ReleaseSection{
					Title:      k,
					Commits:    v,
					IsBreaking: k == breakingChangeTitle,
				})
			}
		}

		c.Releases = append(c.Releases, clr)
	}

	return c, nil
}
