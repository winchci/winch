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

// AssetConfig provides config for asset generate
type AssetConfig struct {
	Enabled   *bool         `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Branches  *FilterConfig `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags      *FilterConfig `json:"tags,omitempty" yaml:"tags,omitempty"`
	Filename  string        `json:"filename,omitempty" yaml:"filename,omitempty"`
	Directory string        `json:"directory,omitempty" yaml:"directory,omitempty"`
	Package   string        `json:"package,omitempty" yaml:"package,omitempty"`
	Variable  string        `json:"variable,omitempty" yaml:"variable,omitempty"`
	Tag       string        `json:"tag,omitempty" yaml:"tag,omitempty"`
	Only      []string      `json:"only,omitempty" yaml:"only,omitempty"`
	Except    []string      `json:"except,omitempty" yaml:"except,omitempty"`
}

func (c *AssetConfig) IsEnabled() bool {
	return c != nil && (c.Enabled == nil || (c.Enabled != nil && *c.Enabled))
}
