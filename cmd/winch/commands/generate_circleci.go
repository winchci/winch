package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/templates"
	"os"
	"path/filepath"
)

type circleCIContext struct {
	Name        string
	Description string
	Repository  string
	Language    string
	Version     string
}

func writeCircleCI(_ context.Context, cfg *config.Config, version string, file string) error {
	f, err := os.Create(filepath.Join(cfg.BasePath, file))
	if err != nil {
		return err
	}

	defer f.Close()

	err = templates.Load(cfg.BasePath, "!circleci.tmpl").Execute(f, &circleCIContext{
		Name:        cfg.Name,
		Description: cfg.Description,
		Repository:  cfg.Repository,
		Language:    cfg.Language,
		Version:     version,
	})
	if err != nil {
		return err
	}

	return nil
}

func generateCircleCI(ctx context.Context) error {
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

	version, _ := getVersionFromReleases(releases)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if len(output) == 0 {
		output = ".circleci/config.yml"
	}

	return writeCircleCI(ctx, cfg, version, output)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "circleci",
		Short: "Generate a CircleCI configuration file",
		Run:   Runner(generateCircleCI),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
