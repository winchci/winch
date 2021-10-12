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
	"github.com/winchci/winch/pkg/config"
	"github.com/winchci/winch/pkg/docker"
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

	dockers := cfg.Dockers
	if cfg.Docker != nil {
		dockers = append(dockers, cfg.Docker)
	}

	for _, dockerConfig := range dockers {
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
		err = d.Build(ctx, cfg, version)
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
