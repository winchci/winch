package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
)

func hookPreCommit(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)
	if c, ok := cfg.Commands["format-check"]; ok {
		err = winch.Run(ctx, c, cfg)
		if err != nil {
			return err
		}
	}

	if c, ok := cfg.Commands["lint"]; ok {
		err = winch.Run(ctx, c, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "pre-commit",
		Short: "Hook for pre-commit",
		Run:   Runner(hookPreCommit),
		Args:  cobra.NoArgs,
	}

	hookCmd.AddCommand(cmd)
}
