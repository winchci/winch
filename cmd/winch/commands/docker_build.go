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

func buildDocker(ctx context.Context, cfg *config.Config) error {
	releases, _, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, _ := getVersionFromReleases(cfg, releases)

	if cfg.Dockerfiles != nil {
		for _, dockerfile := range cfg.Dockerfiles {
			if dockerfile.IsEnabled() {
				err := writeDockerfile(ctx, cfg, dockerfile, version, dockerfile.File)
				if err != nil {
					return err
				}
			}
		}
	}

	contextProvider, err := docker.NewContextProvider()
	if err != nil {
		return err
	}

	defer contextProvider.Close()

	for _, dockerConfig := range append(cfg.Dockers, cfg.Docker) {
		if dockerConfig.IsEnabled() && winch.CheckFilters(ctx, dockerConfig.Branches, dockerConfig.Tags) {
			d, err := docker.NewDocker(dockerConfig, cfg.Name, contextProvider)
			if err != nil {
				return err
			}

			fmt.Printf("Logging into Docker repository %s\n", dockerConfig.Server)
			err = d.Login(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("Building Docker image %s/%s\n", dockerConfig.Organization, dockerConfig.Repository)
			err = d.Build(ctx, cfg, version)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func dockerBuild(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	return buildDocker(ctx, cfg)
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
