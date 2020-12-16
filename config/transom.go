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

// TransomConfig provides configuration for Transom
type TransomConfig struct {
	Enabled      *bool         `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Server       string        `json:"server,omitempty" yaml:"server,omitempty"`
	Organization string        `json:"organization,omitempty" yaml:"organization,omitempty"`
	Application  string        `json:"application,omitempty" yaml:"application,omitempty"`
	Token        string        `json:"token,omitempty" yaml:"token,omitempty"`
	Branches     *FilterConfig `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags         *FilterConfig `json:"tags,omitempty" yaml:"tags,omitempty"`
	Directory    string        `json:"directory,omitempty" yaml:"directory,omitempty"`
	Artifacts    []string      `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

func (d *TransomConfig) IsEnabled() bool {
	return d != nil && (d.Enabled == nil || *d.Enabled) && len(d.Application) > 0 &&
		(len(d.Directory) > 0 || len(d.Artifacts) > 0)
}
