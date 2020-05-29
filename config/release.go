package config

// ReleaseConfig provides config for releases
type ReleaseConfig struct {
	RunConfig
	Artifacts []string      `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

func (d *ReleaseConfig) IsEnabled() bool {
	return d != nil && (d.Enabled == nil || *d.Enabled)
}
