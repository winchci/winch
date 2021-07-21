/*
winch - Universal Build and Release Tool
Copyright (C) 2021 Ketch Kloud, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	winch "github.com/winchci/winch/pkg"
	"github.com/winchci/winch/pkg/config"
	"github.com/winchci/winch/pkg/docker"
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
