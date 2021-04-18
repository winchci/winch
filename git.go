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
	"os"
	"path"
	"strings"

	"github.com/coreos/go-semver/semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
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

func (g Git) GetCommits(_ context.Context) ([]*Commit, error) {
	l, err := g.repo.Log(&git.LogOptions{})
	if err != nil {
		if err.Error() == "reference not found" {
			return nil, nil
		}

		return nil, err
	}

	var commits []*Commit
	err = l.ForEach(func(c *object.Commit) error {
		commits = append(commits, &Commit{
			Hash:    c.Hash.String(),
			When:    c.Author.When,
			Message: ParseMessage(strings.TrimSpace(c.Message)),
		})
		return nil
	})
	if err != nil {
		return nil, err
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
}

func (b Branch) GetName() string {
	return b.name
}

func (b Branch) IsBranch() bool {
	return true
}

func (b Branch) IsTag() bool {
	return false
}

type Tag struct {
	name string
}

func (b Tag) GetName() string {
	return b.name
}

func (b Tag) IsBranch() bool {
	return false
}

func (b Tag) IsTag() bool {
	return true
}

type GitRef interface {
	GetName() string
	IsBranch() bool
	IsTag() bool
}

func (g Git) GetHead(_ context.Context) (GitRef, error) {
	r, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	if r.Name().IsBranch() {
		return &Branch{r.Name().Short()}, nil
	}
	if r.Name().IsTag() {
		return &Tag{r.Name().Short()}, nil
	}

	return nil, fmt.Errorf("unable to determine head ref")
}
