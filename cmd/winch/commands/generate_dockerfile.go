/*
winch - Universal Build and Release Tool
Copyright (C) 2020 Switchbit, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package commands

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
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
