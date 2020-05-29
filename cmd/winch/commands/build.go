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
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
)

func build(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = winch.Run(ctx, cfg.BeforeBuild, cfg)
	if err != nil {
		return err
	}

	err = winch.Run(ctx, cfg.Build, cfg)
	if err != nil {
		return err
	}

	return winch.Run(ctx, cfg.AfterBuild, cfg)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "build",
		Short: "Build",
		Run:   Runner(build),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
