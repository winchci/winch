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

package winch

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/go-semver/semver"
)

type Release struct {
	Version         string
	PreviousVersion string
	Date            time.Time
	Commits         map[CommitType][]*Commit
	FirstCommitHash string
	LastCommitHash  string
	IsNew           bool
}

func MakeReleases(_ context.Context, commits []*Commit, makeRelease bool) ([]*Release, error) {
	var r []*Release
	var current *Release
	var previous *Release
	highestVersion := semver.New("0.0.0")

	for n, c := range commits {
		if n > 0 {
			c.NextHash = commits[n-1].Hash
			commits[n-1].PreviousHash = c.Hash
		}
	}

	for _, c := range commits {
		if current == nil || (current.Version != c.Tag && len(c.Tag) > 0) {
			if current != nil {
				r = append(r, current)
				previous = current
			}

			if len(c.Tag) > 0 && c.Tag[0] == 'v' {
				v := semver.New(c.Tag[1:])
				if highestVersion.LessThan(*v) {
					highestVersion = v
				}
			}

			current = &Release{
				Version:        c.Tag,
				Date:           c.When,
				Commits:        make(map[CommitType][]*Commit),
				LastCommitHash: c.Hash,
			}
			if previous != nil {
				previous.PreviousVersion = current.Version
			}
		}

		current.FirstCommitHash = c.Hash
		current.Commits[c.Message.Type] = append(current.Commits[c.Message.Type], c)
	}

	if current != nil {
		r = append(r, current)
	}

	if len(r) > 0 && r[0].Version == "" {
		if makeRelease {
			bumpedMajor := false
			bumpedMinor := false
			for _, cc := range r[0].Commits {
				for _, c := range cc {
					if c.Message.IsBreaking && !bumpedMajor {
						bumpedMajor = true
					} else if c.Message.Type.IsMinor() {
						bumpedMinor = true
					}
				}
			}

			if bumpedMajor {
				highestVersion.BumpMajor()
			} else if bumpedMinor {
				highestVersion.BumpMinor()
			} else {
				highestVersion.BumpPatch()
			}

			r[0].Version = fmt.Sprintf("v%s", highestVersion.String())
		}
		r[0].IsNew = true
	}

	return r, nil
}
