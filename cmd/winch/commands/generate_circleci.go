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
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
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
	if len(file) == 0 {
		file = ".circleci/config.yml"
	}

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
