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
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type VersionBumpInfo struct {
	Name        string
	Version     string
	Description string
	ReleaseName string
	Prerelease  string
}

func writeVersionFromTemplate(cfg *config.Config, version VersionBumpInfo) error {
	f, err := os.Create(filepath.Join(cfg.BasePath, cfg.Version.File))
	if err != nil {
		return err
	}

	defer f.Close()

	vars := cfg.Version.Variables
	if vars == nil {
		vars = make(map[string]string)
	}
	if _, ok := vars["Name"]; !ok {
		vars["Name"] = version.Name
	}
	if _, ok := vars["Version"]; !ok {
		vars["Version"] = version.Version
	}
	if _, ok := vars["Description"]; !ok {
		vars["Description"] = version.Description
	}
	if _, ok := vars["ReleaseName"]; !ok {
		vars["ReleaseName"] = version.ReleaseName
	}
	if _, ok := vars["Prerelease"]; !ok {
		vars["Prerelease"] = version.Prerelease
	}

	err = templates.Load(cfg.BasePath, cfg.Version.Template).Execute(f, vars)
	if err != nil {
		return err
	}

	return nil
}

func writeVersionNode(cfg *config.Config, version VersionBumpInfo) error {
	b, err := ioutil.ReadFile(cfg.Version.File)
	if err != nil {
		return err
	}

	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	if len(version.Description) > 0 {
		j["description"] = version.Description
	}

	j["version"] = strings.TrimPrefix(version.Version, "v")

	b, err = json.MarshalIndent(j, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cfg.Version.File, b, 0644)
}

func getVersionFromReleases(releases []*winch.Release) (string, string) {
	var version string
	var prerelease string
	if len(releases) > 0 {
		version = releases[0].Version
		prerelease = ""
	} else {
		version = "v0.0.0"
		prerelease = "dev"
	}

	if version[0] == 'v' {
		version = version[1:]
	}

	return version, prerelease
}

func writeVersion(cfg *config.Config, version, prerelease string) error {
	vbi := VersionBumpInfo{
		Name:        cfg.Name,
		Description: cfg.Description,
		Version:     version,
		ReleaseName: winch.Name(context.Background(), "adjectives", "animals"),
		Prerelease:  prerelease,
	}

	if cfg.Language == "node" {
		return writeVersionNode(cfg, vbi)
	} else {
		return writeVersionFromTemplate(cfg, vbi)
	}
}

func generateVersion(ctx context.Context) error {
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

	version, prerelease := getVersionFromReleases(releases)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if len(output) > 0 {
		cfg.Version.File = output
	}

	return writeVersion(cfg, version, prerelease)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Generate the version file",
		Run:   Runner(generateVersion),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
