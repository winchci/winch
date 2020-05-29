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
