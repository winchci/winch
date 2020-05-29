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
	"github.com/winchci/winch/config"
	"os"
	"os/exec"
	"strconv"
)

func dropdb(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = config.LoadDBConfig(ctx, cfg)
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Println("Dropping database", cfg.Database.Database)
	}

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("dropdb",
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		"--if-exists",
		"-i",
		cfg.Database.Database,
	)

	c.Stdin = os.Stdin

	if cfg.Verbose {
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout
	}

	return c.Run()
}

func init() {
	var cmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop the database",
		Run:   Runner(dropdb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())

	dbCmd.AddCommand(cmd)
}
