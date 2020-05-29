package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/version"
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
