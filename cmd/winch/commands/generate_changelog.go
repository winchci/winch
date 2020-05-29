package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/templates"
	"os"
	"path/filepath"
)

func writeChangelog(ctx context.Context, cfg *config.Config, rel []*winch.Release) error {
	cl, err := winch.MakeChangelog(ctx, cfg.Repository, rel)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(cfg.BasePath, cfg.Changelog.File))
	if err != nil {
		return err
	}

	defer f.Close()

	err = templates.Load(cfg.BasePath, cfg.Changelog.Template).Execute(f, cl)
	if err != nil {
		return err
	}

	return nil
}

func makeReleases(ctx context.Context, cfg *config.Config) ([]*winch.Release, error) {
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

func generateChangelog(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cmd := config.CommandFromContext(ctx)

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if len(output) > 0 {
		cfg.Changelog.File = output
	}

	return writeChangelog(ctx, cfg, releases)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "changelog",
		Short: "Generate a changelog",
		Run:   Runner(generateChangelog),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
