package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/docker"
)

func publish(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	return publish2(ctx, cfg)
}

func publish2(ctx context.Context, cfg *config.Config) error {
	shouldPublish := cfg.Publish.IsEnabled()

	if !shouldPublish && cfg.Transom.IsEnabled() && winch.CheckFilters(ctx, cfg.Transom.Branches, cfg.Transom.Tags) {
		shouldPublish = true
	}

	if !shouldPublish {
		for _, dockerConfig := range append(cfg.Dockers, cfg.Docker) {
			if dockerConfig.IsEnabled() && winch.CheckFilters(ctx, dockerConfig.Branches, dockerConfig.Tags) {
				shouldPublish = true
			}
		}
	}

	if !shouldPublish {
		return nil
	}

	fmt.Println("Publishing")

	err := winch.Run(ctx, cfg.BeforePublish, cfg)
	if err != nil {
		return err
	}

	err = winch.Run(ctx, cfg.Publish, cfg)
	if err != nil {
		return err
	}

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, _ := getVersionFromReleases(releases)

	if cfg.Transom.IsEnabled() {
		fmt.Println("Publishing to Transom")
		err = publishToTransom(ctx, cfg, version)
		if err != nil {
			return err
		}
	}

	if cfg.Dockerfile.IsEnabled() {
		err = writeDockerfile(ctx, cfg, version)
		if err != nil {
			return err
		}
	}

	for _, dockerConfig := range append(cfg.Dockers, cfg.Docker) {
		if dockerConfig.IsEnabled() && winch.CheckFilters(ctx, dockerConfig.Branches, dockerConfig.Tags) {
			d, err := docker.NewDocker(dockerConfig, cfg.Name)
			if err != nil {
				return err
			}

			fmt.Printf("Logging into Docker repository %s\n", dockerConfig.Server)
			err = d.Login(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("Building Docker image %s/%s\n", dockerConfig.Organization, dockerConfig.Repository)
			err = d.Build(ctx, version)
			if err != nil {
				return err
			}

			fmt.Printf("Publishing Docker image to %s\n", dockerConfig.Server)
			err = d.Publish(ctx)
			if err != nil {
				return err
			}
		}
	}

	err = winch.Run(ctx, cfg.AfterPublish, cfg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "Publish artifacts to the configured publication points",
		Run:   Runner(publish),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
