package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
)

func test(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = winch.Run(ctx, cfg.BeforeTest, cfg)
	if err != nil {
		return err
	}

	err = winch.Run(ctx, cfg.Test, cfg)
	if err != nil {
		return err
	}

	return winch.Run(ctx, cfg.AfterTest, cfg)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "test",
		Short: "Test",
		Run:   Runner(test),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
