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

package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/pkg/config"
	"github.com/winchci/winch/templates"
	"os"
	"path/filepath"
)

func writeGithubAction(_ context.Context, cfg *config.Config, t *config.TemplateFileConfig, version string, file string) error {
	if !t.IsEnabled() {
		return nil
	}

	if len(file) == 0 {
		file = t.GetFile()
	}

	if len(file) == 0 {
		file = ".github/workflows/ci.yml"
	}

	f, err := os.Create(filepath.Join(cfg.BasePath, file))
	if err != nil {
		return err
	}

	defer f.Close()

	vars := t.GetVariables()
	if vars == nil {
		vars = make(map[string]string)
	}
	if _, ok := vars["Name"]; !ok {
		vars["Name"] = cfg.Name
	}
	if _, ok := vars["Description"]; !ok {
		vars["Description"] = cfg.Description
	}
	if _, ok := vars["Repository"]; !ok {
		vars["Repository"] = cfg.Repository
	}
	if _, ok := vars["Language"]; !ok {
		vars["Language"] = cfg.Language
	}
	if _, ok := vars["Version"]; !ok {
		vars["Version"] = version
	}

	err = templates.Load(cfg.BasePath, t.GetTemplate()).Execute(f, vars)
	if err != nil {
		return err
	}

	return nil
}

func generateGithubAction(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cmd := config.CommandFromContext(ctx)

	releases, _, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, _ := getVersionFromReleases(cfg, releases)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	return writeGithubAction(ctx, cfg, cfg.GitHubAction, version, output)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "action",
		Short: "Generate a GitHub Action configuration file",
		Run:   Runner(generateGithubAction),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
