package config

// TransomConfig provides configuration for Transom
type TransomConfig struct {
	Enabled      *bool         `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Server       string        `json:"server,omitempty" yaml:"server,omitempty"`
	Shipyard     string        `json:"shipyard,omitempty" yaml:"shipyard,omitempty"`
	Organization string        `json:"organization,omitempty" yaml:"organization,omitempty"`
	Application  string        `json:"application,omitempty" yaml:"application,omitempty"`
	Token        string        `json:"token,omitempty" yaml:"token,omitempty"`
	Username     string        `json:"username,omitempty" yaml:"username,omitempty"`
	Password     string        `json:"password,omitempty" yaml:"password,omitempty"`
	Branches     *FilterConfig `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags         *FilterConfig `json:"tags,omitempty" yaml:"tags,omitempty"`
	Directory    string        `json:"directory,omitempty" yaml:"directory,omitempty"`
	Artifacts    []string      `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

func (d *TransomConfig) IsEnabled() bool {
	return d != nil && (d.Enabled == nil || *d.Enabled) && len(d.Application) > 0 &&
		(len(d.Directory) > 0 || len(d.Artifacts) > 0)
}
