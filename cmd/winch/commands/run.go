package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
)

func run(ctx context.Context, args []string) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cmd := cfg.Commands[args[0]]
	if cmd == nil {
		return fmt.Errorf("command '%s' not found", args[0])
	}

	return winch.Run(ctx, cmd, cfg)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "run COMMAND",
		Short: "Run a command",
		Run:   RunnerWithArgs(run),
		Args:  cobra.ExactArgs(1),
	}

	rootCmd.AddCommand(cmd)
}
