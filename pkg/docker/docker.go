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

package docker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/iancoleman/strcase"
	"github.com/winchci/winch/pkg"
	"github.com/winchci/winch/pkg/config"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Docker struct {
	cfg  *config.DockerConfig
	name string
}

func NewDocker(cfg *config.DockerConfig, name string) (*Docker, error) {
	if cfg.Context == "" {
		cfg.Context = "."
	}

	if len(cfg.Server) == 0 {
		cfg.Server = os.Getenv("DOCKERHUB_SERVER")
	}

	if len(cfg.Server) == 0 {
		cfg.Server = "docker.io"
	}

	if len(cfg.Organization) == 0 {
		cfg.Organization = os.Getenv("DOCKERHUB_ORGANIZATION")
	}

	if len(cfg.Organization) == 0 {
		return nil, fmt.Errorf("the DockerHub organization is required")
	}

	if len(cfg.Repository) == 0 {
		cfg.Repository = os.Getenv("DOCKERHUB_REPOSITORY")
	}

	if len(cfg.Repository) == 0 {
		cfg.Repository = name
	}

	if len(cfg.Repository) == 0 {
		return nil, fmt.Errorf("the DockerHub repository is required")
	}

	if len(cfg.Username) == 0 {
		cfg.Username = os.Getenv("DOCKERHUB_USERNAME")
	}

	if len(cfg.Username) == 0 {
		return nil, fmt.Errorf("the DockerHub usename is required")
	}

	if len(cfg.Password) == 0 {
		cfg.Password = os.Getenv("DOCKERHUB_PASSWORD")
	}

	if len(cfg.Password) == 0 {
		return nil, fmt.Errorf("the DockerHub password is required")
	}

	if len(cfg.Tag) == 0 {
		cfg.Tag = "latest"
	}

	cfg.Server = os.ExpandEnv(cfg.Server)
	cfg.Organization = os.ExpandEnv(cfg.Organization)
	cfg.Repository = os.ExpandEnv(cfg.Repository)
	cfg.Username = os.ExpandEnv(cfg.Username)
	cfg.Password = os.ExpandEnv(cfg.Password)

	return &Docker{
		cfg: cfg,
	}, nil
}

func (d *Docker) Close(ctx context.Context) error {
	if len(d.name) > 0 {
		args := []string{"docker", "buildx", "rm", "--force", d.name}
		fmt.Println(strings.Join(args, " "))
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}

		d.name = ""
	}

	return nil
}

func (d *Docker) Login(ctx context.Context) error {
	args := []string{"docker", "login", "--username", d.cfg.Username, "--password-stdin", d.cfg.Server}
	fmt.Println(strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = bytes.NewReader([]byte(d.cfg.Password))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (d *Docker) Build(ctx context.Context, cfg *config.Config, version string) error {
	return d.build(ctx, cfg, version, false)
}

func (d *Docker) build(ctx context.Context, cfg *config.Config, version string, push bool) error {
	d.name = strcase.ToSnake(pkg.Name(ctx, "adjectives", "animals"))

	args := []string{"docker", "buildx", "create", "--name", d.name, "--use", "--buildkitd-flags", "--allow-insecure-entitlement network.host", "--driver", "docker-container"}
	fmt.Println(strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	v, err := semver.NewVersion(version)
	if err != nil {
		return err
	}

	var tags []string
	baseTag := fmt.Sprintf("%s/%s/%s", d.cfg.Server, d.cfg.Organization, d.cfg.Repository)
	if d.cfg.Tag == "latest" {
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, d.cfg.Tag))
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, version))
		tags = append(tags, fmt.Sprintf("%s:%d", baseTag, v.Major))
		tags = append(tags, fmt.Sprintf("%s:%d.%d", baseTag, v.Major, v.Minor))
	} else if len(d.cfg.Tag) > 0 {
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, d.cfg.Tag))
		tags = append(tags, fmt.Sprintf("%s:%s-%s", baseTag, version, d.cfg.Tag))
		tags = append(tags, fmt.Sprintf("%s:%d-%s", baseTag, v.Major, d.cfg.Tag))
		tags = append(tags, fmt.Sprintf("%s:%d.%d-%s", baseTag, v.Major, v.Minor, d.cfg.Tag))
	} else {
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, version))
		tags = append(tags, fmt.Sprintf("%s:%d", baseTag, v.Major))
		tags = append(tags, fmt.Sprintf("%s:%d.%d", baseTag, v.Major, v.Minor))
	}

	if d.cfg.Labels == nil {
		d.cfg.Labels = make(map[string]string)
	}
	d.cfg.Labels["org.opencontainers.image.source"] = cfg.Repository
	d.cfg.Labels["org.opencontainers.image.created"] = time.Now().UTC().Format(time.RFC3339)
	d.cfg.Labels["org.opencontainers.image.version"] = version
	d.cfg.Labels["org.opencontainers.image.title"] = cfg.Name
	d.cfg.Labels["org.opencontainers.image.description"] = cfg.Description

	args = []string{"docker", "buildx", "build", "--file", d.cfg.Dockerfile, "--allow", "network.host", "--builder", d.name}
	if push {
		args = append(args, "--push")
		if len(d.cfg.Platforms) > 0 {
			args = append(args, "--platform", strings.Join(d.cfg.Platforms, ","))
		}
	} else {
		args = append(args, "--load")
	}

	for _, tag := range tags {
		args = append(args, "--tag", tag)
	}

	for k, v := range d.cfg.Labels {
		args = append(args, "--label", fmt.Sprintf("%s=%s", k, v))
	}

	for k, v := range d.cfg.BuildArgs {
		args = append(args, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	if len(d.cfg.Target) > 0 {
		args = append(args, "--target", d.cfg.Target)
	}

	args = append(args, d.cfg.Context)
	fmt.Println(strings.Join(args, " "))
	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	if !push {
		image := fmt.Sprintf("%s/%s/%s", d.cfg.Server, d.cfg.Organization, d.cfg.Repository)
		snykAuthToken := os.Getenv("SNYK_AUTH_TOKEN")

		if (d.cfg.Scan == nil || *d.cfg.Scan) && len(snykAuthToken) > 0 {
			args := []string{"docker", "scan", "--accept-license", "--login"}
			if len(snykAuthToken) > 0 {
				fmt.Println(strings.Join(append(args, "--token", "*****"), " "))

				args = append(args, "--token", snykAuthToken)
			} else {
				fmt.Println(strings.Join(args, " "))
			}

			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				return err
			}

			args = []string{"docker", "scan", "--accept-license", "--severity", "medium", image}
			fmt.Println(strings.Join(args, " "))
			cmd = exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err = cmd.Run()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Docker) Publish(ctx context.Context, cfg *config.Config, version string) error {
	return d.build(ctx, cfg, version, true)
}
