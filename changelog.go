package winch

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
