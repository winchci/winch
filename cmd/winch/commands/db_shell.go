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
	"github.com/winchci/winch/config"
	"os"
	"os/exec"
	"strconv"
)

func shelldb(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = config.LoadDBConfig(ctx, cfg)
	if err != nil {
		return err
	}

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("psql",
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		cfg.Database.Database,
	)

	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}

func init() {
	var cmd = &cobra.Command{
		Use:   "shell",
		Short: "Open a database shell",
		Run:   Runner(shelldb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())

	dbCmd.AddCommand(cmd)
}
