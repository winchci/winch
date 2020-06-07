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
	"github.com/winchci/winch/changelog"
	"github.com/winchci/winch/config"
)

func generateVersion(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cmd2 := config.CommandFromContext(ctx)

	releases, err := changelog.MakeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, prerelease := changelog.GetVersionFromReleases(releases)

	output, err := cmd2.Flags().GetString("output")
	if err != nil {
		return err
	}

	if len(output) > 0 {
		cfg.Version.File = output
	}

	return changelog.WriteVersion(cfg, version, prerelease)
}

func init() {
	var cmd2 = &cobra.Command{
		Use:   "version",
		Short: "Generate the version file",
		Run:   Runner(generateVersion),
		Args:  cobra.NoArgs,
	}

	cmd2.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd2)
}
