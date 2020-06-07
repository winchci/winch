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
	config2 "github.com/winchci/winch/v2/actions"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func upgrade(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cfg2, err := config2.Upgrade(cfg)
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(cfg2)
	if err != nil {
		return err
	}

	cmd2 := config.CommandFromContext(ctx)

	output, err := cmd2.Flags().GetString("file")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(output, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd2 = &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the configuration",
		Run:   Runner(upgrade),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd2)
}
