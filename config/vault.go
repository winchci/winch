package config

// The keys for the parameters
type VaultConfig struct {
	Address string `json:"address" yaml:"address"`
	Token   string `json:"token" yaml:"token"`
	Prefix  string `json:"prefix" yaml:"prefix"`
	Dir     string `json:"dir" yaml:"dir"`
}
