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
)

func up(ctx context.Context) error {
	return fmt.Errorf("not implemented (yet)")
}

func init() {
	var cmd = &cobra.Command{
		Use:   "up",
		Short: "Bring up an environment",
		Run:   Runner(up),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringSlice("except", nil, "image to exclude")
	cmd.Flags().StringSlice("only", nil, "image to spin up")

	rootCmd.AddCommand(cmd)
}
