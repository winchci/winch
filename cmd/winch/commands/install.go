package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
)

func install(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = winch.Run(ctx, cfg.BeforeInstall, cfg)
	if err != nil {
		return err
	}

	err = winch.Run(ctx, cfg.Install, cfg)
	if err != nil {
		return err
	}

	return winch.Run(ctx, cfg.AfterInstall, cfg)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "install",
		Short: "Install",
		Run:   Runner(install),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
