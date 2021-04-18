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

package config

// DockerConfig provides config for Docker
type DockerConfig struct {
	Enabled      *bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Server       string             `json:"server,omitempty" yaml:"server,omitempty"`
	Organization string             `json:"organization,omitempty" yaml:"organization,omitempty"`
	Repository   string             `json:"repository,omitempty" yaml:"repository,omitempty"`
	Username     string             `json:"username,omitempty" yaml:"username,omitempty"`
	Password     string             `json:"password,omitempty" yaml:"password,omitempty"`
	Dockerfile   string             `json:"dockerfile,omitempty" yaml:"dockerfile,omitempty"`
	Context      string             `json:"context,omitempty" yaml:"context,omitempty"`
	Tag          string             `json:"tag,omitempty" yaml:"tag,omitempty"`
	BuildArgs    map[string]*string `json:"buildargs,omitempty" yaml:"buildargs,omitempty"`
	Labels       map[string]string  `json:"labels,omitempty" yaml:"labels,omitempty"`
	Branches     *FilterConfig      `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags         *FilterConfig      `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func (d *DockerConfig) IsEnabled() bool {
	return d != nil && (d.Enabled == nil || *d.Enabled) && len(d.Organization) > 0
}
