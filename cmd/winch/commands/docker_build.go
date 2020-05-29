package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/docker"
)

func dockerBuild(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, _ := getVersionFromReleases(releases)

	for _, dockerConfig := range append(cfg.Dockers, cfg.Docker) {
		d, err := docker.NewDocker(dockerConfig, cfg.Name)
		if err != nil {
			return err
		}

		fmt.Println("Logging into DockerHub")
		err = d.Login(ctx)
		if err != nil {
			return err
		}

		fmt.Print("Building Docker image")
		err = d.Build(ctx, version)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "build",
		Short: "Build docker image",
		Run:   Runner(dockerBuild),
		Args:  cobra.NoArgs,
	}

	dockerCmd.AddCommand(cmd)
}
