package changelog

import (
	"context"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
)

func MakeReleases(ctx context.Context, cfg *config.Config) ([]*winch.Release, error) {
	var g winch.Repository
	var err error

	if cfg.Local {
		g, err = winch.NewGit(ctx, winch.FindGitDir(ctx))
	} else {
		g, err = winch.NewGitHub(ctx, cfg.Repository)
	}
	if err != nil {
		return nil, err
	}

	tags, err := g.GetTags(ctx)
	if err != nil {
		return nil, err
	}

	commits, err := g.GetCommits(ctx)
	if err != nil {
		return nil, err
	}

	winch.TagCommits(ctx, commits, tags)

	releases, err := winch.MakeReleases(ctx, commits, true)
	if err != nil {
		return nil, err
	}

	return releases, err
}
