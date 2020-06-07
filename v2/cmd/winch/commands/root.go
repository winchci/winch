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
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/v2/project"
	"github.com/winchci/winch/v2/version"
	"log"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:               version.Name,
	Short:             version.Description,
	Version:           version.String(),
	PersistentPreRunE: setup,
	Run:               runner(root),
	Args:              cobra.NoArgs,
	TraverseChildren:  true,
}

// Setup sets up the configuration system.
func setup(cmd *cobra.Command, args []string) error {
	_ = godotenv.Overload()
	return nil
}

func root(ctx context.Context, c *cobra.Command, args []string) error {
	file, err := c.Flags().GetString("file")
	if err != nil {
		return err
	}

	p, err := project.LoadProject(ctx, file)
	if err != nil {
		return err
	}

	pm := project.NewManager()

	err = pm.Execute(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runner(f func(context.Context, *cobra.Command, []string) error) func(*cobra.Command, []string) {
	return func(c *cobra.Command, args []string) {
		err := f(context.Background(), c, args)
		if err != nil {
			parts := make([]string, 0)
			for c != nil {
				parts = append([]string{c.Name()}, parts...)
				c = c.Parent()
			}
			name := strings.Join(parts, " ")
			fmt.Printf("%s: %s\n", name, err)
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

	rootCmd.PersistentFlags().StringP("file", "f", "winch.yml", "configuration file")
}
