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
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"io/ioutil"
)

func generateEnvironment(ctx context.Context, args []string) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	if env, ok := cfg.Environments[args[0]]; ok {
		buf := bytes.NewBuffer(nil)
		for k, v := range env {
			buf.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		}
		return ioutil.WriteFile(".env", buf.Bytes(), 0644)
	} else {
		return fmt.Errorf("environment %s not defined", args[0])
	}
}

func init() {
	var cmd = &cobra.Command{
		Use:   "environment",
		Short: "Generate .env file",
		Run:   RunnerWithArgs(generateEnvironment),
		Args:  cobra.ExactArgs(1),
	}

	generateCmd.AddCommand(cmd)
}
