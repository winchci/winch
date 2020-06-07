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
	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/vfsgen"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func generateAsset(ctx context.Context, asset *config.AssetConfig) error {
	if !asset.IsEnabled() {
		return nil
	}

	if !winch.CheckFilters(ctx, asset.Branches, asset.Tags) {
		return nil
	}

	var d http.FileSystem
	d = http.Dir(asset.Directory)

	if len(asset.Only) != 0 {
		d = filter.Keep(d, func(path string, fi os.FileInfo) bool {
			if fi.IsDir() {
				return true
			}

			for _, p := range asset.Only {
				if ok, err := regexp.MatchString(p, path); err == nil && ok {
					return true
				}
			}

			return false
		})
	}

	if len(asset.Except) != 0 {
		d = filter.Skip(d, func(path string, fi os.FileInfo) bool {
			if fi.IsDir() {
				return false
			}

			for _, p := range asset.Except {
				if ok, err := filepath.Match(p, path); err == nil && ok {
					return true
				}
			}

			return false
		})
	}

	err := vfsgen.Generate(d, vfsgen.Options{
		Filename:     asset.Filename,
		PackageName:  asset.Package,
		BuildTags:    asset.Tag,
		VariableName: asset.Variable,
	})

	return err
}

func generateAssets(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	for _, asset := range cfg.Assets {
		err = generateAsset(ctx, asset)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd2 = &cobra.Command{
		Use:   "assets",
		Short: "Generate assets",
		Run:   Runner(generateAssets),
		Args:  cobra.NoArgs,
	}

	generateCmd.AddCommand(cmd2)
}
