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
)

func recreateVault(ctx context.Context) error {
	err := dropVault(ctx)
	if err != nil {
		return err
	}

	err = createVault(ctx)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "recreate",
		Short: "Recreate the database",
		Run:   Runner(recreateVault),
		Args:  cobra.NoArgs,
	}

	vaultCmd.AddCommand(cmd)
}
