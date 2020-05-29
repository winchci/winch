package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/docker"
)

func dockerPublish(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	for _, dockerConfig := range append(cfg.Dockers, cfg.Docker) {
		if !winch.CheckFilters(ctx, dockerConfig.Branches, dockerConfig.Tags) {
			return nil
		}

		d, err := docker.NewDocker(dockerConfig, cfg.Name)
		if err != nil {
			return err
		}

		fmt.Printf("Logging into Docker repository %s\n", dockerConfig.Server)
		err = d.Login(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Publishing Docker image to %s\n", dockerConfig.Server)
		err = d.Publish(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "Publish container to DockerHub",
		Run:   Runner(dockerPublish),
		Args:  cobra.NoArgs,
	}

	dockerCmd.AddCommand(cmd)
}
