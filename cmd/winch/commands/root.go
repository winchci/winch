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
	"github.com/winchci/winch/version"
	"log"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:               version.Name,
	Short:             version.Description,
	Version:           version.String(),
	TraverseChildren:  true,
	PersistentPreRunE: config.Setup(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func commandName(c *cobra.Command) string {
	parts := make([]string, 0)
	for c != nil {
		parts = append([]string{c.Name()}, parts...)
		c = c.Parent()
	}
	return strings.Join(parts[1:], " ")
}

func Runner(f func(ctx context.Context) error) func(c *cobra.Command, args []string) {
	return func(c *cobra.Command, args []string) {
		err := f(config.AddCommandToContext(context.Background(), c))
		if err != nil {
			name := commandName(c)
			fmt.Printf("%s: failed\n%s", name, err)
			os.Exit(1)
		}
	}
}

func RunnerWithArgs(f func(ctx context.Context, args []string) error) func(c *cobra.Command, args []string) {
	return func(c *cobra.Command, args []string) {
		err := f(config.AddCommandToContext(context.Background(), c), args)
		if err != nil {
			name := commandName(c)
			fmt.Printf("%s: failed\n%s", name, err)
			os.Exit(1)
		}
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help [command]",
		Short: "Help about any command",
		Long: `Help provides help for any command in the application.
Simply type ` + rootCmd.Name() + ` help [path to command] for full details.`,

		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Printf("Unknown help topic %#q\n", args)
				c.Root().Usage()
			} else {
				cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
				cmd.Help()
			}
		},
	})

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet output")
	rootCmd.PersistentFlags().StringP("file", "f", "winch.yml", "configuration file")
}
