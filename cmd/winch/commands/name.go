package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
)

func name(ctx context.Context, args []string) error {
	fmt.Println(winch.Name(ctx, args...))
	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "name",
		Short: "Generate a release name",
		Run:   RunnerWithArgs(name),
		Args:  cobra.MinimumNArgs(1),
	}

	rootCmd.AddCommand(cmd)
}
