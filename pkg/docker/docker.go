package docker

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/winchci/winch/pkg/config"
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
	cfg             *config.DockerConfig
	client          *client.Client
	contextProvider *ContextProvider
	token           string
}

func NewDocker(cfg *config.DockerConfig, name string, contextProvider *ContextProvider) (*Docker, error) {
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

	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &Docker{
		cfg:             cfg,
		client:          c,
		contextProvider: contextProvider,
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

func (d *Docker) Build(ctx context.Context, cfg *config.Config, version string) error {
	contextArchive, err := d.contextProvider.GetContext(d.cfg.Context)
	if err != nil {
		return err
	}

	f, err := os.Open(contextArchive)
	if err != nil {
		return err
	}

	defer f.Close()

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

	c, err := d.loadConfig()
	if err != nil {
		return err
	}

	resp, err := d.client.ImageBuild(ctx, f, types.ImageBuildOptions{
		Tags:        tags,
		Dockerfile:  d.cfg.Dockerfile,
		BuildArgs:   d.cfg.BuildArgs,
		Labels:      d.cfg.Labels,
		Target:      d.cfg.Target,
		AuthConfigs: c.Auths,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return printDockerResponse(resp.Body)
}

func (d *Docker) Publish(ctx context.Context) error {
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

type dockerConfig struct {
	Auths map[string]types.AuthConfig `json:"auths"`
}

func (d *Docker) loadConfig() (*dockerConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filename := filepath.Join(home, ".docker", "config.json")

	f, err := os.Open(filename)
	if err != nil {
		filename = filepath.Join("/etc", "docker", "config.json")

		f, _ = os.Open(filename)
	}

	cfg := &dockerConfig{
		Auths: make(map[string]types.AuthConfig),
	}

	if f != nil {
		defer f.Close()
		err = json.NewDecoder(f).Decode(cfg)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range cfg.Auths {
		if len(v.Auth) > 0 {
			b, err := base64.StdEncoding.DecodeString(v.Auth)
			if err != nil {
				return nil, err
			}

			parts := strings.Split(string(b), ":")
			if len(parts) >= 2 {
				cfg.Auths[k] = types.AuthConfig{
					Username: parts[0],
					Password: parts[1],
				}
			}
		}
	}

	cfg.Auths[d.cfg.Server] = types.AuthConfig{
		Username:      d.cfg.Username,
		Password:      d.cfg.Password,
		ServerAddress: d.cfg.Server,
	}

	return cfg, nil
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
