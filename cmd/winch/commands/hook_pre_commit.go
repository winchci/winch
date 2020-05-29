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

func hookPreCommit(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)
	if c, ok := cfg.Commands["format-check"]; ok {
		err = winch.Run(ctx, c, cfg)
		if err != nil {
			return err
		}
	}

	if c, ok := cfg.Commands["lint"]; ok {
		err = winch.Run(ctx, c, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "pre-commit",
		Short: "Hook for pre-commit",
		Run:   Runner(hookPreCommit),
		Args:  cobra.NoArgs,
	}

	hookCmd.AddCommand(cmd)
}
