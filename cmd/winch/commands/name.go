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
	"fmt"
	"github.com/spf13/cobra"
	winch "github.com/winchci/winch/pkg"
)

func name(ctx context.Context, args []string) error {
	fmt.Println(winch.Name(ctx, args...))
	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "name",
		Short: "Generate a release name",
		Run:   RunnerWithArgs(name),
		Args:  cobra.MinimumNArgs(1),
	}

	rootCmd.AddCommand(cmd)
}
