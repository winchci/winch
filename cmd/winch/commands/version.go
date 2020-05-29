package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
)

func showVersion(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, prerelease := getVersionFromReleases(releases)

	if len(prerelease) > 0 {
		fmt.Printf("%s-%s\n", version, prerelease)
	} else {
		fmt.Println(version)
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Show the version",
		Run:   Runner(showVersion),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
