package project

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func LoadProject(ctx context.Context, filename string) (*Project, error) {
	if len(filename) == 0 {
		filename = "winch.yml"
	}

	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("cannot open project file %s", filename)
	}

	b , err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := new(Project)
	err = yaml.Unmarshal(b, cfg)
	if err != nil {
		return nil, err
	}

	for _, s := range cfg.Steps {
		s.Project = cfg
		s.Environment = copyEnv(s.Environment, cfg.Environment)
		if s.Container == nil {
			s.Container = cfg.Container
		}
	}

	for k, j := range cfg.Jobs {
		j.ID = k

		if j.Container == nil {
			j.Container = cfg.Container
		}
		j.Environment = copyEnv(j.Environment, cfg.Environment)

		for _, s := range j.Steps {
			if len(s.Shell) > 0 {
				s.Shell = j.Shell
			}
			if s.Container == nil {
				s.Container = j.Container
			}

			s.Environment = copyEnv(s.Environment, j.Environment)
			s.Job = j
			s.Project = cfg
		}
	}

	return cfg, nil
}

func copyEnv(dst map[string]string, src map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string)
	}
	for k, v := range src {
		if _, ok := dst[k]; !ok {
			dst[k] = v
		}
	}
	return dst
}
