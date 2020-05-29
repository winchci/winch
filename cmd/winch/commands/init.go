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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type PackageJsonRepository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type PackageJson struct {
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Repository      *PackageJsonRepository `json:"repository"`
	Scripts         map[string]string      `json:"scripts"`
	Dependencies    map[string]string      `json:"dependencies"`
	DevDependencies map[string]string      `json:"devDependencies"`
}

func exist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func initialize(ctx context.Context) error {
	_, err := config.LoadConfig(ctx)
	if err == nil {
		return fmt.Errorf("this directory already has a winch configuration")
	}

	cmd := config.CommandFromContext(ctx)

	c := &config.Config{}

	if exist("package.json") {
		b, err := ioutil.ReadFile("package.json")
		if err == nil {
			var p PackageJson
			err = json.Unmarshal(b, &p)
			if err != nil {
				return err
			}

			c.Name = p.Name
			c.Description = p.Description

			if p.Repository != nil && p.Repository.Type == "git" {
				c.Repository = p.Repository.Url
			}

			if len(c.Repository) == 0 {
				c.Repository = "TODO"
			}

			c.Language = "node"
			c.Toolchain = "npm"

			// Try to detect Yarn
			for k := range p.DevDependencies {
				if strings.Contains(k, "yarn") {
					c.Toolchain = "yarn"
				}
			}

			for k := range p.Dependencies {
				if strings.Contains(k, "yarn") {
					c.Toolchain = "yarn"
				}
			}

			for k := range p.Scripts {
				if strings.Contains(k, "yarn") {
					c.Toolchain = "yarn"
				}
			}
		}
	} else {
		m, err := winch.LoadGoModuleDefinition(ctx, ".")
		if err != nil {
			return err
		}

		c.Repository = fmt.Sprintf("https://%s", m.Module.Path)
		parts := strings.Split(c.Repository, "/")
		c.Name = parts[len(parts)-1]
	}

	c.Release = &config.ReleaseConfig{}
	c.Release.Branches = &config.FilterConfig{
		Only: "master",
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	filename, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a configuration file",
		Run:   Runner(initialize),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
