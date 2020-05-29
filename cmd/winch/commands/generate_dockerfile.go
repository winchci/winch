package commands

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/templates"
	"os"
	"path/filepath"
)

type dockerfileContext struct {
	Name        string
	Description string
	Language    string
	Version     string
}

func writeDockerfile(_ context.Context, cfg *config.Config, version string) error {
	if !cfg.Dockerfile.IsEnabled() {
		return nil
	}

	f, err := os.Create(filepath.Join(cfg.BasePath, cfg.Dockerfile.File))
	if err != nil {
		return err
	}

	defer f.Close()

	err = templates.Load(cfg.BasePath, cfg.Dockerfile.Template).Execute(f, &dockerfileContext{
		Name:        cfg.Name,
		Description: cfg.Description,
		Language:    cfg.Language,
		Version:     version,
	})
	if err != nil {
		return err
	}

	return nil
}

func generateDockerfile(ctx context.Context) error {
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

	cfg.Dockerfile.Enabled = proto.Bool(true)
	if len(cfg.Dockerfile.File) == 0 {
		cfg.Dockerfile.File = "Dockerfile"
	}

	if len(output) > 0 {
		cfg.Dockerfile.File = output
	}

	return writeDockerfile(ctx, cfg, version)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "dockerfile",
		Short: "Generate a Dockerfile",
		Run:   Runner(generateDockerfile),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
