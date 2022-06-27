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

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// FilterConfig provides config for filters
type FilterConfig struct {
	Ignore string `json:"ignore,omitempty" yaml:"ignore,omitempty"`
	Only   string `json:"only,omitempty" yaml:"only,omitempty"`
}

func (c *FilterConfig) IsEnabled() bool {
	return c != nil && (len(c.Ignore) > 0 || len(c.Only) > 0)
}

// CIConfig provides config for CI
type CIConfig struct {
	Enabled  *bool         `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Branches *FilterConfig `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags     *FilterConfig `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func (c *CIConfig) IsEnabled() bool {
	return c != nil && (c.Enabled == nil || (c.Enabled != nil && *c.Enabled))
}

// RunConfig provides config for commands to run
type RunConfig struct {
	Enabled     *bool             `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Command     string            `json:"command,omitempty" yaml:"command,omitempty"`
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Branches    *FilterConfig     `json:"branches,omitempty" yaml:"branches,omitempty"`
	Tags        *FilterConfig     `json:"tags,omitempty" yaml:"tags,omitempty"`
	Input       string            `json:"input,omitempty" yaml:"input,omitempty"`
}

func (c *RunConfig) IsEnabled() bool {
	return c != nil && (c.Enabled == nil || (c.Enabled != nil && *c.Enabled)) && len(c.Command) > 0
}

// TemplateFileConfig provides config for files produced from templates
type TemplateFileConfig struct {
	Enabled   *bool             `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Template  string            `json:"template,omitempty" yaml:"template,omitempty"`
	File      string            `json:"file,omitempty" yaml:"file,omitempty"`
	Variables map[string]string `json:"variables,omitempty" yaml:"variables,omitempty"`
}

func (c TemplateFileConfig) GetFile() string {
	return c.File
}

func (c TemplateFileConfig) GetTemplate() string {
	return c.Template
}

func (c *TemplateFileConfig) IsEnabled() bool {
	return c != nil && ((c.Enabled == nil && len(c.File) > 0) || (c.Enabled != nil && *c.Enabled))
}

func (c TemplateFileConfig) GetVariables() map[string]string {
	return c.Variables
}

// Config provides overall configuration
type Config struct {
	Filename      string                       `json:"filename,omitempty" yaml:"filename,omitempty"`
	BasePath      string                       `json:"base,omitempty" yaml:"base,omitempty"`
	Name          string                       `json:"name,omitempty" yaml:"name,omitempty"`
	Description   string                       `json:"description,omitempty" yaml:"description,omitempty"`
	Repository    string                       `json:"repository,omitempty" yaml:"repository,omitempty"`
	Local         bool                         `json:"local,omitempty" yaml:"local,omitempty"`
	Prerelease    string                       `json:"prerelease,omitempty" yaml:"prerelease,omitempty"`
	Verbose       bool                         `json:"verbose,omitempty" yaml:"verbose,omitempty"`
	Quiet         bool                         `json:"quiet,omitempty" yaml:"quiet,omitempty"`
	Mono          bool                         `json:"mono,omitempty" yaml:"mono,omitempty"`
	Parallelism   int                          `json:"arallelism,omitempty" yaml:"parallelism,omitempty"`
	Language      string                       `json:"language,omitempty" yaml:"language,omitempty"`
	Toolchain     string                       `json:"toolchain,omitempty" yaml:"toolchain,omitempty"`
	CI            *CIConfig                    `json:"ci,omitempty" yaml:"ci,omitempty"`
	BeforeInstall *RunConfig                   `json:"before_install,omitempty" yaml:"before_install,omitempty"`
	Install       *RunConfig                   `json:"install,omitempty" yaml:"install,omitempty"`
	AfterInstall  *RunConfig                   `json:"after_install,omitempty" yaml:"after_install,omitempty"`
	BeforeBuild   *RunConfig                   `json:"before_build,omitempty" yaml:"before_build,omitempty"`
	Build         *RunConfig                   `json:"build,omitempty" yaml:"build,omitempty"`
	AfterBuild    *RunConfig                   `json:"after_build,omitempty" yaml:"after_build,omitempty"`
	BeforeTest    *RunConfig                   `json:"before_test,omitempty" yaml:"before_test,omitempty"`
	Test          *RunConfig                   `json:"test,omitempty" yaml:"test,omitempty"`
	AfterTest     *RunConfig                   `json:"after_test,omitempty" yaml:"after_test,omitempty"`
	Changelog     *TemplateFileConfig          `json:"changelog,omitempty" yaml:"changelog,omitempty"`
	BeforeVersion *RunConfig                   `json:"before_version,omitempty" yaml:"before_version,omitempty"`
	Version       *TemplateFileConfig          `json:"version,omitempty" yaml:"version,omitempty"`
	Versions      []*TemplateFileConfig        `json:"versions,omitempty" yaml:"versions,omitempty"`
	AfterVersion  *RunConfig                   `json:"after_version,omitempty" yaml:"after_version,omitempty"`
	GitHubAction  *TemplateFileConfig          `json:"githubaction,omitempty" yaml:"githubaction,omitempty"`
	Dockerfile    *TemplateFileConfig          `json:"dockerfile,omitempty" yaml:"dockerfile,omitempty"`
	Dockerfiles   []*TemplateFileConfig        `json:"dockerfiles,omitempty" yaml:"dockerfiles,omitempty"`
	Homebrew      *HomebrewConfig              `json:"homebrew,omitempty" yaml:"homebrew,omitempty"`
	Transom       *TransomConfig               `json:"transom,omitempty" yaml:"transom,omitempty"`
	Docker        *DockerConfig                `json:"docker,omitempty" yaml:"docker,omitempty"`
	Dockers       []*DockerConfig              `json:"dockers,omitempty" yaml:"dockers,omitempty"`
	Assets        []*AssetConfig               `json:"assets,omitempty" yaml:"assets,omitempty"`
	BeforeRelease *RunConfig                   `json:"before_release,omitempty" yaml:"before_release,omitempty"`
	Release       *ReleaseConfig               `json:"release,omitempty" yaml:"release,omitempty"`
	AfterRelease  *RunConfig                   `json:"after_release,omitempty" yaml:"after_release,omitempty"`
	BeforePublish *RunConfig                   `json:"before_publish,omitempty" yaml:"before_publish,omitempty"`
	Publish       *RunConfig                   `json:"publish,omitempty" yaml:"publish,omitempty"`
	AfterPublish  *RunConfig                   `json:"after_publish,omitempty" yaml:"after_publish,omitempty"`
	Environment   map[string]string            `json:"environment,omitempty" yaml:"environment,omitempty"`
	Environments  map[string]map[string]string `json:"environments,omitempty" yaml:"environments,omitempty"`
	Scopes        []string                     `json:"scopes,omitempty" yaml:"scopes,omitempty"`
	Artifacts     []string                     `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

func makeBool(b bool) *bool {
	return &b
}

// Default configurations
var (
	homebrew = &HomebrewConfig{
		Template: "!brew.tmpl",
		File:     "formula.rb",
	}

	DefaultMonoConfig = &Config{
		Language: "mono",
		Install: &RunConfig{
			Enabled: makeBool(false),
		},
		Build: &RunConfig{
			Enabled: makeBool(false),
		},
		Test: &RunConfig{
			Enabled: makeBool(false),
		},
		Version: &TemplateFileConfig{
			Enabled: makeBool(false),
		},
		GitHubAction: &TemplateFileConfig{
			Enabled: makeBool(false),
		},
		Homebrew: homebrew,
	}

	DefaultDockerConfig = &Config{
		Language: "docker",
		Install: &RunConfig{
			Enabled: makeBool(false),
		},
		Build: &RunConfig{
			Enabled: makeBool(false),
		},
		Test: &RunConfig{
			Enabled: makeBool(false),
		},
		Version: &TemplateFileConfig{
			Enabled: makeBool(false),
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!docker_action.tmpl",
		},
		Homebrew: homebrew,
	}

	DefaultGoConfig = &Config{
		Language: "go",
		Install: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Download modules",
			Command: "go mod download",
		},
		Build: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Build",
			Command: "go build ./...",
		},
		Test: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Test",
			Command: "go test ./...",
		},
		Version: &TemplateFileConfig{
			File:     "version/version_gen.go",
			Template: "!version_go.tmpl",
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!go_action.tmpl",
		},
		Dockerfile: &TemplateFileConfig{
			Template: "!go_dockerfile.tmpl",
		},
		Homebrew: homebrew,
	}

	DefaultNpmConfig = &Config{
		Language:  "node",
		Toolchain: "npm",
		Install: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Download modules",
			Command: "npm install",
		},
		Build: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Build",
			Command: "npm run build",
		},
		Test: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Test",
			Command: "npm test",
		},
		Version: &TemplateFileConfig{
			File: "package.json",
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!node_action.tmpl",
		},
		Dockerfile: &TemplateFileConfig{
			Template: "!node_npm_dockerfile.tmpl",
		},
		Homebrew: homebrew,
	}

	DefaultYarnConfig = &Config{
		Language:  "node",
		Toolchain: "yarn",
		Install: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Download modules",
			Command: "yarn",
		},
		Build: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Build",
			Command: "yarn run build",
		},
		Test: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Test",
			Command: "yarn test",
		},
		Version: &TemplateFileConfig{
			Enabled: makeBool(true),
			File:    "package.json",
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!node_action.tmpl",
		},
		Dockerfile: &TemplateFileConfig{
			Template: "!node_yarn_dockerfile.tmpl",
		},
		Homebrew: homebrew,
	}

	DefaultHelmConfig = &Config{
		Language: "helm",
		Install: &RunConfig{
			Enabled: makeBool(false),
			Name:    "Download modules",
			//Command: "yarn",
		},
		Build: &RunConfig{
			Enabled: makeBool(false),
			Name:    "Build",
			//Command: "yarn run build",
		},
		Test: &RunConfig{
			Enabled: makeBool(false),
			Name:    "Test",
			//Command: "yarn test",
		},
		Version: &TemplateFileConfig{
			Enabled: makeBool(false),
			//File:    "package.json",
		},
	}

	DefaultJavaMavenConfig = &Config{
		Language:  "java",
		Toolchain: "mvn",
		Install: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Validate",
			Command: "mvn -B -Dbuild.version=${BUILD_VERSION} clean process-resources",
		},
		Build: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Build",
			Command: "mvn -B -Dbuild.version=${BUILD_VERSION} compile",
		},
		Test: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Test",
			Command: "mvn -B -Dbuild.version=${BUILD_VERSION} test",
		},
		Publish: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Publish",
			Command: "mvn -B -Dbuild.version=${BUILD_VERSION} deploy",
		},
		Version: &TemplateFileConfig{
			Enabled: makeBool(false),
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!java_mvn_action.tmpl",
		},
		Dockerfile: &TemplateFileConfig{
			Template: "!java_mvn_dockerfile.tmpl",
		},
		Homebrew: homebrew,
	}

	DefaultScalaSbtConfig = &Config{
		Language:  "scala",
		Toolchain: "sbt",
		Install: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Download modules",
			Command: "sbt install",
		},
		Build: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Build",
			Command: "sbt compile",
		},
		Test: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Test",
			Command: "sbt test",
		},
		Version: &TemplateFileConfig{
			Enabled: makeBool(false),
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!scala_sbt_action.tmpl",
		},
		Dockerfile: &TemplateFileConfig{
			Template: "!scala_sbt_dockerfile.tmpl",
		},
		Homebrew: homebrew,
	}

	DefaultPythonConfig = &Config{
		Language: "python",
		Install: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Download modules",
			Command: "pip install .",
		},
		Build: &RunConfig{
			Enabled: makeBool(false),
		},
		Test: &RunConfig{
			Enabled: makeBool(true),
			Name:    "Run tests",
			Command: "python -m unittest -v",
		},
		Version: &TemplateFileConfig{
			File:     "version.py",
			Template: "!version_py.tmpl",
		},
		GitHubAction: &TemplateFileConfig{
			Enabled:  makeBool(true),
			Template: "!python_action.tmpl",
		},
		Dockerfile: &TemplateFileConfig{
			Template: "!python_dockerfile.tmpl",
		},
		Homebrew: homebrew,
	}
)

type commandContext struct{}
type configContext struct{}

func CommandFromContext(ctx context.Context) *cobra.Command {
	v := ctx.Value(&commandContext{})
	if v == nil {
		return nil
	}
	return v.(*cobra.Command)
}

func AddCommandToContext(ctx context.Context, c *cobra.Command) context.Context {
	return context.WithValue(ctx, &commandContext{}, c)
}

func ConfigFromContext(ctx context.Context) *Config {
	v := ctx.Value(&configContext{})
	if v == nil {
		return nil
	}
	return v.(*Config)
}

func AddConfigToContext(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, &configContext{}, c)
}

// LoadConfig loads config from the configuration file
func LoadConfig(ctx context.Context) (context.Context, error) {
	c := CommandFromContext(ctx)

	filename := c.Flags().Lookup("file").Value.String()
	for _, f := range strings.Split(filename, ":") {
		f, err := filepath.Abs(f)
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(f); err == nil {
			filename = f
		}
	}

	cfg := &Config{
		Filename: filename,
		BasePath: filepath.Dir(filename),
		Language: "go",
		Changelog: &TemplateFileConfig{
			Template: "!changelog.tmpl",
			File:     "CHANGELOG.md",
		},
	}

	b, err := ioutil.ReadFile(cfg.Filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, cfg)
	if err != nil {
		return nil, err
	}

	verbose, err := c.Flags().GetBool("verbose")
	if err != nil {
		return nil, err
	}

	if verbose {
		cfg.Verbose = true
	}

	quiet, err := c.Flags().GetBool("quiet")
	if err != nil {
		return nil, err
	}

	if quiet {
		cfg.Quiet = true
	}

	if cfg.Language == "" {
		cfg.Language, cfg.Toolchain = GuessLanguageAndToolchain()
	}

	var defaultConfig *Config
	switch cfg.Language {
	case "node":
		switch cfg.Toolchain {
		case "yarn":
			defaultConfig = DefaultYarnConfig
		default:
			defaultConfig = DefaultNpmConfig
		}

	case "helm":
		defaultConfig = DefaultHelmConfig

	case "java":
		defaultConfig = DefaultJavaMavenConfig

	case "scala":
		defaultConfig = DefaultScalaSbtConfig

	case "python":
		defaultConfig = DefaultPythonConfig
		if fileExists("requirements.txt") {
			defaultConfig.Install.Enabled = proto.Bool(true)
			defaultConfig.Install.Command = "pip install -r requirements.txt"
		} else if fileExists("requirements.pip") {
			defaultConfig.Install.Enabled = proto.Bool(true)
			defaultConfig.Install.Command = "pip install -r requirements.pip"
		}

	case "docker":
		defaultConfig = DefaultDockerConfig

	default:
		if cfg.Mono {
			defaultConfig = DefaultMonoConfig
		} else {
			defaultConfig = DefaultGoConfig
			if dirExists("cmd") && fileExists(fmt.Sprintf("cmd/%s/main.go", cfg.Name)) {
				defaultConfig.Build.Command = fmt.Sprintf("go build -o bin/%s ./cmd/%s/main.go", cfg.Name, cfg.Name)
			}
		}
	}

	if cfg.Install == nil {
		cfg.Install = defaultConfig.Install
	}
	if cfg.Install.Enabled == nil {
		cfg.Install.Enabled = makeBool(true)
	}

	if cfg.Build == nil {
		cfg.Build = defaultConfig.Build
	}
	if cfg.Build.Enabled == nil {
		cfg.Build.Enabled = makeBool(true)
	}

	if cfg.Test == nil {
		cfg.Test = defaultConfig.Test
	}
	if cfg.Test.Enabled == nil {
		cfg.Test.Enabled = makeBool(true)
	}

	if cfg.Publish == nil {
		cfg.Publish = defaultConfig.Publish
	}
	if cfg.Publish != nil && cfg.Publish.Enabled == nil {
		cfg.Publish.Enabled = makeBool(true)
	}

	if cfg.GitHubAction == nil {
		cfg.GitHubAction = defaultConfig.GitHubAction
	} else {
		if len(cfg.GitHubAction.Template) == 0 {
			cfg.GitHubAction.Template = defaultConfig.GitHubAction.Template
		}

		if len(cfg.GitHubAction.File) == 0 {
			cfg.GitHubAction.File = defaultConfig.GitHubAction.File
		}
	}

	if cfg.Version == nil {
		cfg.Version = defaultConfig.Version
	} else {
		if len(cfg.Version.Template) == 0 {
			cfg.Version.Template = defaultConfig.Version.Template
		}

		if len(cfg.Version.File) == 0 {
			cfg.Version.File = defaultConfig.Version.File
		}
	}

	if cfg.Dockerfile == nil {
		cfg.Dockerfile = defaultConfig.Dockerfile
	} else {
		if len(cfg.Dockerfile.Template) == 0 {
			cfg.Dockerfile.Template = defaultConfig.Dockerfile.Template
		}

		if len(cfg.Dockerfile.File) == 0 {
			cfg.Dockerfile.File = defaultConfig.Dockerfile.File
		}
	}

	if cfg.Homebrew == nil {
		cfg.Homebrew = defaultConfig.Homebrew
	} else {
		if len(cfg.Homebrew.File) == 0 {
			cfg.Homebrew.File = defaultConfig.Homebrew.File
		}

		if len(cfg.Homebrew.Template) == 0 {
			cfg.Homebrew.Template = defaultConfig.Homebrew.Template
		}

		if len(cfg.Homebrew.Install) == 0 {
			cfg.Homebrew.Install = fmt.Sprintf("bin.install \"%s\"", cfg.Name)
		}

		if len(cfg.Homebrew.Test) == 0 {
			cfg.Homebrew.Test = fmt.Sprintf("system \"#{bin}/%s --version\"", cfg.Name)
		}
	}

	// Force git if a monorepo
	if cfg.Mono {
		cfg.Local = true
	}

	return AddConfigToContext(ctx, cfg), nil
}

// Setup sets up the configuration system.
func Setup() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		_ = godotenv.Overload()
		return nil
	}
}
