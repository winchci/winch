package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"os"
	"strings"
)

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
