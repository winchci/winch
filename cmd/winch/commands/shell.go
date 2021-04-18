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
	"github.com/winchci/winch/config"
	"os"
	"os/exec"
)

func startShell(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)
	cmd := config.CommandFromContext(ctx)

	sh := os.Getenv("SHELL")
	if len(sh) == 0 {
		sh = "/bin/sh"
	}

	c := exec.CommandContext(ctx, sh)
	c.Env = os.Environ()

	for k, v := range cfg.Environment {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}

	if envName, err := cmd.Flags().GetString("env"); err != nil && len(envName) > 0 {
		if e, ok := cfg.Environments[envName]; ok {
			for k, v := range e {
				c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
			}
		} else {
			return fmt.Errorf("environment %s not defined", envName)
		}
	} else if err != nil {
		return err
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}

func init() {
	var cmd = &cobra.Command{
		Use:   "shell",
		Short: "Open a shell",
		Run:   Runner(startShell),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().String("env", "", "environment to load")

	rootCmd.AddCommand(cmd)
}
