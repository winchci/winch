package config

// HomebrewConfig provides config for Homebrew
type HomebrewConfig struct {
	Enabled      *bool                  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Organization string                 `json:"organization,omitempty" yaml:"organization,omitempty"`
	Repository   string                 `json:"repository,omitempty" yaml:"repository,omitempty"`
	Template     string                 `json:"template,omitempty" yaml:"template,omitempty"`
	File         string                 `json:"file,omitempty" yaml:"file,omitempty"`
	Url          string                 `json:"url,omitempty" yaml:"url,omitempty"`
	Asset        string                 `json:"asset,omitempty" yaml:"asset,omitempty"`
	Install      string                 `json:"install,omitempty" yaml:"install,omitempty"`
	Test         string                 `json:"test,omitempty" yaml:"test,omitempty"`
	DependsOn    []string               `json:"depends_on,omitempty" yaml:"depends_on,omitempty"`
	Variables    map[string]interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`
	Branches     *FilterConfig          `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags         *FilterConfig          `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func (c *HomebrewConfig) GetFile() string {
	return c.File
}

func (c *HomebrewConfig) GetTemplate() string {
	return c.Template
}

func (c *HomebrewConfig) IsEnabled() bool {
	return c != nil && ((c.Enabled == nil && len(c.File) > 0) || (c.Enabled != nil && *c.Enabled)) && len(c.Organization) > 0
}
