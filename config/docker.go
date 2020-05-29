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
	Branches     *FilterConfig      `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags         *FilterConfig      `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func (d *DockerConfig) IsEnabled() bool {
	return d != nil && (d.Enabled == nil || *d.Enabled) && len(d.Organization) > 0
}
