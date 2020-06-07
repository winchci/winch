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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func writeHomebrew(ctx context.Context, cfg *config.Config, t *config.HomebrewConfig, version string, file string) error {
	if !t.IsEnabled() {
		return nil
	}

	if len(file) == 0 {
		file = t.GetFile()
	}

	if len(file) == 0 {
		file = "formula.rb"
	}

	f, err := os.Create(filepath.Join(cfg.BasePath, file))
	if err != nil {
		return err
	}

	defer f.Close()

	vars := t.Variables
	if vars == nil {
		vars = make(map[string]interface{})
	}
	if _, ok := vars["Name"]; !ok {
		vars["Name"] = cfg.Name
	}
	if _, ok := vars["Description"]; !ok {
		vars["Description"] = cfg.Description
	}
	if _, ok := vars["Repository"]; !ok {
		vars["Repository"] = fmt.Sprintf("homebrew-%s", vars["Name"])
	}
	if _, ok := vars["Language"]; !ok {
		vars["Language"] = cfg.Language
	}
	if _, ok := vars["Homepage"]; !ok {
		vars["Homepage"] = cfg.Repository
	}
	if _, ok := vars["Version"]; !ok {
		vars["Version"] = version
	}
	if _, ok := vars["Install"]; !ok {
		vars["Install"] = t.Install
	}
	if _, ok := vars["Test"]; !ok {
		vars["Test"] = t.Test
	}
	if _, ok := vars["DependsOn"]; !ok {
		vars["DependsOn"] = t.DependsOn
	}
	if _, ok := vars["Asset"]; !ok {
		vars["Asset"] = t.Asset
	}
	if _, ok := vars["Url"]; !ok && len(t.Url) > 0 {
		vars["Url"] = t.Url
	}
	if _, ok := vars["Url"]; !ok {
		vars["Url"] = fmt.Sprintf("%s/releases/download/v%s/%s", cfg.Repository, vars["Version"], vars["Asset"])
	}

	var data []byte
	if strings.HasPrefix(vars["Url"].(string), "http") {
		fmt.Printf("homebrew: downloading '%s'\n", vars["Url"].(string))

		req, err := http.NewRequestWithContext(ctx, "GET", vars["Url"].(string), nil)
		req.SetBasicAuth("sethyates", os.Getenv("GITHUB_TOKEN"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check server response
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("homebrew: bad status: %s", resp.Status)
		}

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("homebrew: opening file '%s'\n", vars["Url"].(string))
		data, err = ioutil.ReadFile(vars["Url"].(string))
		if err != nil {
			return err
		}
	}

	bytes := sha256.Sum256(data)
	sha := hex.EncodeToString(bytes[:])
	vars["Sha256"] = sha

	err = templates.Load(cfg.BasePath, t.GetTemplate()).Execute(f, vars)
	if err != nil {
		return err
	}

	return nil
}

func generateHomebrew(ctx context.Context) error {
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

	return writeHomebrew(ctx, cfg, cfg.Homebrew, version, output)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "homebrew",
		Short: "Generate a Homebrew formula",
		Run:   Runner(generateHomebrew),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
