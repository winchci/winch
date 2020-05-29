package winch

import (
	"context"
	"time"
)

type Commit struct {
	Hash         string
	PreviousHash string
	NextHash     string
	When         time.Time
	Message      *Message
	Tag          string
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
