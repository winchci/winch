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
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mholt/archiver/v3"
	"github.com/winchci/winch/pkg/config"
	"github.com/winchci/winch/version"
)

type dockerErrorMsg struct {
	Message string `json:"message"`
}

type dockerMsg struct {
	Status      string         `json:"status"`
	Message     string         `json:"stream"`
	ErrorDetail dockerErrorMsg `json:"errorDetail"`
}

type Docker struct {
	cfg    *config.DockerConfig
	client *client.Client
	token  string
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

	c, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Docker{
		cfg:    cfg,
		client: c,
	}, nil
}

func (d *Docker) Login(ctx context.Context) error {
	_, err := d.client.RegistryLogin(ctx, types.AuthConfig{
		Username:      d.cfg.Username,
		Password:      d.cfg.Password,
		ServerAddress: d.cfg.Server,
	})
	if err != nil {
		return err
	}

	b, err := json.Marshal(types.AuthConfig{
		Username:      d.cfg.Username,
		Password:      d.cfg.Password,
		ServerAddress: d.cfg.Server,
	})
	if err != nil {
		return err
	}

	d.token = base64.StdEncoding.EncodeToString(b)

	return nil
}

func (d Docker) Build(ctx context.Context, tag string) error {
	dir, err := ioutil.TempDir("", version.Name)
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	context := path.Join(dir, "context.tar")
	err = archiver.Archive([]string{d.cfg.Context}, context)
	if err != nil {
		return err
	}

	f, err := os.Open(context)
	if err != nil {
		return err
	}

	defer f.Close()

	var tags []string
	baseTag := fmt.Sprintf("%s/%s/%s", d.cfg.Server, d.cfg.Organization, d.cfg.Repository)
	if d.cfg.Tag == "latest" {
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, d.cfg.Tag))
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, tag))
	} else if len(d.cfg.Tag) > 0 {
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, d.cfg.Tag))
		tags = append(tags, fmt.Sprintf("%s:%s-%s", baseTag, tag, d.cfg.Tag))
	} else {
		tags = append(tags, fmt.Sprintf("%s:%s", baseTag, tag))
	}

	resp, err := d.client.ImageBuild(ctx, f, types.ImageBuildOptions{
		Tags:       tags,
		Dockerfile: d.cfg.Dockerfile,
		BuildArgs:  d.cfg.BuildArgs,
		Labels:     d.cfg.Labels,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return printDockerResponse(resp.Body)
}

func (d Docker) Publish(ctx context.Context) error {
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

	resp, err := d.client.ImagePush(ctx, image, types.ImagePushOptions{
		RegistryAuth: d.token,
		All:          true,
	})
	if err != nil {
		return err
	}
	defer resp.Close()

	return printDockerResponse(resp)
}

func printDockerResponse(body io.ReadCloser) error {
	r := bufio.NewReader(body)
	for {
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		var msg dockerMsg
		err = json.Unmarshal(line, &msg)
		if err != nil {
			return err
		}

		if len(msg.Message) > 0 {
			os.Stdout.WriteString(strings.TrimSpace(msg.Message))
			os.Stdout.WriteString("\n")
		}
		if len(msg.Status) > 0 {
			os.Stdout.WriteString(strings.TrimSpace(msg.Status))
			os.Stdout.WriteString("\n")
		}
		if len(msg.ErrorDetail.Message) > 0 {
			return errors.New(msg.ErrorDetail.Message)
		}
	}

	return nil
}
