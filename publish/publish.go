/*
winch - Universal Build and Release Tool
Copyright (C) 2020 Switchbit, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package publish

import (
	"context"
	"fmt"
	"github.com/winchci/winch"
	"github.com/winchci/winch/changelog"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/docker"
	"github.com/winchci/winch/homebrew"
	"github.com/winchci/winch/transom"
)

func DoPublish(ctx context.Context, cfg *config.Config) error {
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

	releases, err := changelog.MakeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, _ := changelog.GetVersionFromReleases(releases)

	if cfg.Transom.IsEnabled() {
		fmt.Println("Publishing to Transom")
		err = transom.PublishToTransom(ctx, cfg, version)
		if err != nil {
			return err
		}
	}

	if cfg.Homebrew.IsEnabled() {
		err = homebrew.WriteHomebrew(ctx, cfg, cfg.Homebrew, version, cfg.Homebrew.File)
		if err != nil {
			return err
		}
	}

	if cfg.Dockerfile.IsEnabled() {
		err = docker.WriteDockerfile(ctx, cfg, cfg.Dockerfile, version, cfg.Dockerfile.File)
		if err != nil {
			return err
		}
	}

	if cfg.Dockerfiles != nil {
		for _, dockerfile := range cfg.Dockerfiles {
			if dockerfile.IsEnabled() {
				err = docker.WriteDockerfile(ctx, cfg, dockerfile, version, dockerfile.File)
				if err != nil {
					return err
				}
			}
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
