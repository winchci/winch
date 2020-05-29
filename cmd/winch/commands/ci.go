package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"os"
)

func ci(ctx context.Context) error {
	if os.Getenv("CI") != "true" {
		return fmt.Errorf("must be running in a CI environment")
	}

	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	// CI is always verbose and never quiet
	cfg.Verbose = true
	cfg.Quiet = false

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, prerelease := getVersionFromReleases(releases)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Prerelease: %s\n", prerelease)

	err = os.Setenv("BUILD_VERSION", version)
	if err != nil {
		return err
	}

	if cfg.Version.IsEnabled() {
		fmt.Println("Creating version")
		err = writeVersion(cfg, version, prerelease)
		if err != nil {
			return err
		}
	}

	if cfg.Changelog.IsEnabled() {
		fmt.Println("Creating changelog")
		err = writeChangelog(ctx, cfg, releases)
		if err != nil {
			return err
		}
	}

	if cfg.Install.IsEnabled() {
		fmt.Println("Installing")

		err = winch.Run(ctx, cfg.BeforeInstall, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.Install, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.AfterInstall, cfg)
		if err != nil {
			return err
		}
	}

	fmt.Println("Creating assets")
	for _, asset := range cfg.Assets {
		err = generateAsset(ctx, asset)
		if err != nil {
			return err
		}
	}

	if cfg.Build.IsEnabled() {
		fmt.Println("Building")

		err = winch.Run(ctx, cfg.BeforeBuild, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.Build, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.AfterBuild, cfg)
		if err != nil {
			return err
		}
	}

	if cfg.Test.IsEnabled() {
		fmt.Println("Testing")

		err = winch.Run(ctx, cfg.BeforeTest, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.Test, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.AfterTest, cfg)
		if err != nil {
			return err
		}
	}

	if cfg.Release.IsEnabled() {
		fmt.Println("Releasing")
		err = release2(ctx, cfg)
		if err != nil {
			return err
		}
	}

	err = publish2(ctx, cfg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "ci",
		Short: "Execute a CI build",
		Run:   Runner(ci),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
