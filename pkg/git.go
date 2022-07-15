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
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/winchci/winch/pkg/config"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Git struct {
	repo *git.Repository
}

func NewGit(_ context.Context, dir string) (*Git, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, err
	}

	return &Git{
		repo: repo,
	}, nil
}

func (g Git) GetTags(_ context.Context) (map[string]string, error) {
	tagrefs, err := g.repo.Tags()
	if err != nil {
		return nil, err
	}

	tags := map[string]string{}

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		s := t.Name().Short()
		h := t.Hash().String()
		if s[0] == 'v' {
			newVersion := semver.New(s[1:])
			if _, ok := tags[h]; !ok {
				tags[h] = s
			} else {
				existingVersion := semver.New(tags[h][1:])
				if existingVersion.LessThan(*newVersion) {
					tags[h] = s
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (g Git) GetCommits(ctx context.Context) ([]*Commit, error) {
	l, err := g.repo.Log(&git.LogOptions{})
	if err != nil {
		if err.Error() == "reference not found" {
			return nil, nil
		}

		return nil, err
	}

	cfg := config.ConfigFromContext(ctx)

	var commits []*Commit
	for i := 0; i < cfg.CommitLength; i++ {
		c, err := l.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		affectedPaths := make(map[string]bool)

		if cfg.Mono {
			stats, err := c.Stats()
			if err != nil {
				return nil, err
			}

			for _, stat := range stats {
				if stat.Addition > 0 || stat.Deletion > 0 {
					path := strings.Split(stat.Name, "/")
					depth := len(path)
					if len(path) > cfg.MonoDepth {
						depth = cfg.MonoDepth
					}
					affectedPaths[filepath.Join(path[0:depth]...)] = true
				}
			}
		}

		commit := &Commit{
			Hash:    c.Hash.String(),
			When:    c.Author.When,
			Message: ParseMessage(strings.TrimSpace(c.Message)),
		}

		for affectedPath := range affectedPaths {
			commit.AffectedPaths = append(commit.AffectedPaths, affectedPath)
		}

		commits = append(commits, commit)

		if cfg.Verbose {
			fmt.Printf("Parsed through commit %s\n", commit.Hash)
		}
	}

	return commits, nil
}

func FindGitDir(_ context.Context) string {
	cwd, err := os.Getwd()
	if err == nil {
		for {
			if i, err := os.Stat(path.Join(cwd, ".git")); err == nil && i.IsDir() {
				return cwd
			}

			cwd = path.Join(cwd, "..")
			if cwd == "/" {
				break
			}
		}
	}

	return "."
}

type Branch struct {
	name string
	hash string
}

func (b Branch) GetName() string {
	return b.name
}

func (b Branch) GetHash() string {
	return b.hash
}

func (b Branch) IsBranch() bool {
	return true
}

func (b Branch) IsTag() bool {
	return false
}

type Tag struct {
	name string
	hash string
}

func (b Tag) GetName() string {
	return b.name
}

func (b Tag) GetHash() string {
	return b.hash
}

func (b Tag) IsBranch() bool {
	return false
}

func (b Tag) IsTag() bool {
	return true
}

type GitRef interface {
	GetName() string
	GetHash() string
	IsBranch() bool
	IsTag() bool
}

func (g Git) GetHead(_ context.Context) (GitRef, error) {
	r, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	if r.Name().IsBranch() {
		return &Branch{
			name: r.Name().Short(),
			hash: r.Hash().String(),
		}, nil
	}
	if r.Name().IsTag() {
		return &Tag{
			name: r.Name().Short(),
			hash: r.Hash().String(),
		}, nil
	}

	return nil, fmt.Errorf("unable to determine head ref")
}
