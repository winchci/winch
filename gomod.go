package winch

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
)

type Module struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

type GoMod struct {
	Module  Module    `json:"module"`
	Go      string    `json:"go"`
	Require []Require `json:"require"`
	Exclude []Module  `json:"exclude"`
	Replace []Replace `json:"replace"`
}

type Require struct {
	Path     string `json:"path"`
	Version  string `json:"version"`
	Indirect bool   `json:"indirect"`
}

type Replace struct {
	Old Module `json:"old"`
	New Module `json:"new"`
}

// LoadGoModuleDefinition loads the go.mod definition in a given directory
func LoadGoModuleDefinition(ctx context.Context, dir string) (*GoMod, error) {
	c := exec.CommandContext(ctx, "sh", "-c", "go mod edit -json")
	c.Dir = dir

	buf := new(bytes.Buffer)

	c.Stdout = buf
	err := c.Run()
	if err != nil {
		return nil, err
	}

	var m GoMod
	err = json.Unmarshal(buf.Bytes(), &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
