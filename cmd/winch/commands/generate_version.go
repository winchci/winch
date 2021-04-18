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
	"encoding/xml"
	"github.com/spf13/cobra"
	"github.com/vifraa/gopom"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
	"gopkg.in/yaml.v3"
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

func writeVersionToFile(cfg *config.Config, file *config.TemplateFileConfig, version VersionBumpInfo) error {
	if !file.IsEnabled() {
		return nil
	}

	filename := filepath.Join(cfg.BasePath, file.File)

	if filepath.Base(filename) == "VERSION" && len(file.Template) == 0 {
		file.Template = "!VERSION.tmpl"
	}

	if filepath.Base(filename) == "pom.xml" {
		pom, err := gopom.Parse(filename)
		if err != nil {
			return err
		}

		pom.Version = version.Version
		b, err := xml.MarshalIndent(pom, "", "\t")
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(filename, b, 0644); err != nil {
			return err
		}

		return nil
	}

	switch filepath.Ext(filename) {
	case ".json":
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			b = []byte("{}")
		}

		var j map[string]interface{}
		err = json.Unmarshal(b, &j)
		if err != nil {
			return err
		}

		path := []string{"version"}

		if file.Variables != nil && len(file.Variables["path"]) > 0 {
			path = strings.Split(file.Variables["path"], ".")
		}

		v := j
		for _, key := range path[0:len(path)-1] {
			if _, ok := v[key]; !ok {
				v[key] = make(map[string]interface{})
			}

			v = v[key].(map[string]interface{})
		}

		v[path[len(path)-1]] = strings.TrimPrefix(version.Version, "v")

		b, err = json.MarshalIndent(j, "", "\t")
		if err != nil {
			return err
		}

		return ioutil.WriteFile(filename, b, 0644)

	case ".yaml", ".yml":
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			b = []byte("version:\n")
		}

		var j map[string]interface{}
		err = yaml.Unmarshal(b, &j)
		if err != nil {
			return err
		}

		path := []string{"version"}

		if file.Variables != nil && len(file.Variables["path"]) > 0 {
			path = strings.Split(file.Variables["path"], ".")
		}

		v := j
		for _, key := range path[0:len(path)-1] {
			if _, ok := v[key]; !ok {
				v[key] = make(map[string]interface{})
			}

			v = v[key].(map[string]interface{})
		}

		v[path[len(path)-1]] = strings.TrimPrefix(version.Version, "v")

		b, err = yaml.Marshal(j)
		if err != nil {
			return err
		}

		return ioutil.WriteFile(filename, b, 0644)

	default:
		err := os.MkdirAll(filepath.Dir(filename), 0750)
		if err != nil {
			return err
		}

		f, err := os.Create(filename)
		if err != nil {
			return err
		}

		defer f.Close()

		vars := file.Variables
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

		err = templates.Load(cfg.BasePath, file.Template).Execute(f, vars)
		if err != nil {
			return err
		}

		return nil
	}
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

	err := writeVersionToFile(cfg, cfg.Version, vbi)
	if err != nil {
		return err
	}

	for _, file := range cfg.Versions {
		err := writeVersionToFile(cfg, file, vbi)
		if err != nil {
			return err
		}
	}

	return nil
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
