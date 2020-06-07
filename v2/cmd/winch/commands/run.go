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
	"github.com/winchci/winch/v2/project"
)

func run(ctx context.Context, c *cobra.Command, args []string) error {
	file, err := c.Flags().GetString("file")
	if err != nil {
		return err
	}

	p, err := project.LoadProject(ctx, file)
	if err != nil {
		return err
	}

	pm := project.NewManager()

	err = pm.ExecuteJob(ctx, p, args[0])
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "run",
		Short: "Run a job",
		Run:   runner(run),
		Args:  cobra.ExactArgs(1),
	}

	rootCmd.AddCommand(cmd)
}
