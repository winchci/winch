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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
)

func run(ctx context.Context, args []string) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cmd := cfg.Commands[args[0]]
	if cmd == nil {
		return fmt.Errorf("command '%s' not found", args[0])
	}

	return winch.Run(ctx, cmd, cfg)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "run COMMAND",
		Short: "Run a command",
		Run:   RunnerWithArgs(run),
		Args:  cobra.ExactArgs(1),
	}

	rootCmd.AddCommand(cmd)
}
