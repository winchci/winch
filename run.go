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

package winch

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/winchci/winch/config"
)

func Run(ctx context.Context, cmd *config.RunConfig, cfg *config.Config) error {
	if !cmd.IsEnabled() {
		return nil
	}

	if !CheckFilters(ctx, cmd.Branches, cmd.Tags) {
		return nil
	}

	c := exec.CommandContext(ctx, "sh", "-c", cmd.Command)
	env := make(map[string]string)

	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		env[parts[0]] = parts[1]
	}

	for k, v := range cfg.Environment {
		env[k] = v
	}

	for k, v := range cmd.Environment {
		env[k] = v
	}

	for replacementsRequired := true; replacementsRequired; {
		replacementsRequired = false

		for k, v := range env {
			env[k] = os.Expand(v, func(s string) string {
				return env[s]
			})

			if strings.Contains(env[k], "$") {
				replacementsRequired = true
			}
		}
	}

	for k, v := range env {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}

	if len(cmd.Input) > 0 {
		c.Stdin = bytes.NewBufferString(cmd.Input)
	}

	if cfg.Verbose {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	if len(cmd.Name) > 0 {
		fmt.Println(cmd.Name)
	}

	if !cfg.Quiet {
		fmt.Println(cmd.Command)
	}

	return c.Run()
}
