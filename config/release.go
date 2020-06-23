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

package config

// ReleaseConfig provides config for releases
type ReleaseConfig struct {
	Enabled     *bool             `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Command     string            `json:"command,omitempty" yaml:"command,omitempty"`
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Branches    *FilterConfig     `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags        *FilterConfig     `json:"tags,omitempty" yaml:"tags,omitempty"`
	Input       string            `json:"input,omitempty" yaml:"input,omitempty"`
	Artifacts   []string          `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

func (c *ReleaseConfig) IsEnabled() bool {
	return c != nil && (c.Enabled == nil || (c.Enabled != nil && *c.Enabled)) && len(c.Command) > 0
}

func (c *ReleaseConfig) RunConfig() *RunConfig {
	if c == nil {
		return nil
	}

	return &RunConfig{
		Enabled:     c.Enabled,
		Name:        c.Name,
		Command:     c.Command,
		Environment: c.Environment,
		Branches:    c.Branches,
		Tags:        c.Tags,
		Input:       c.Input,
	}
}
