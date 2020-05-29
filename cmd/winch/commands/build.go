package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
)

func build(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = winch.Run(ctx, cfg.BeforeBuild, cfg)
	if err != nil {
		return err
	}

	err = winch.Run(ctx, cfg.Build, cfg)
	if err != nil {
		return err
	}

	return winch.Run(ctx, cfg.AfterBuild, cfg)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "build",
		Short: "Build",
		Run:   Runner(build),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
