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

package pkg

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
