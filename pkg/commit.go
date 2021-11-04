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

type Commit struct {
	Hash          string
	PreviousHash  string
	NextHash      string
	When          time.Time
	Message       *Message
	Tag           string
	AffectedPaths []string
}

func (c Commit) ShortHash() string {
	return c.Hash[0:8]
}

func (c Commit) Title() string {
	return c.Message.Title()
}

func TagCommits(_ context.Context, commits []*Commit, tags map[string]string) {
	var lastTag string
	for _, c := range commits {
		c.Tag = tags[c.Hash]
		if len(c.Tag) == 0 {
			c.Tag = lastTag
		}
		lastTag = c.Tag
	}
}
