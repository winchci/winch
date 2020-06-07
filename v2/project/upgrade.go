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

package project

import (
	"fmt"
	"github.com/winchci/winch/config"
)

func upgradeFilterConfig(filter *config.FilterConfig) *Filter {
	if filter == nil {
		return nil
	}

	return &Filter{
		Ignore: filter.Ignore,
		Only:   filter.Only,
	}
}

type HasEnabled interface {
	IsEnabled() bool
}

func upgradeCondition(e HasEnabled) string {
	if e != nil && !e.IsEnabled() {
		return "false"
	}
	return ""
}

func upgradeRunConfig(run *config.RunConfig) *Step {
	if run == nil {
		return nil
	}

	var condition string
	if run.Enabled != nil {
		condition = upgradeCondition(run)
	}

	return &Step{
		Name:        run.Name,
		Run:         run.Command,
		Environment: run.Environment,
		If:          condition,
		Branches:    upgradeFilterConfig(run.Branches),
		Tags:        upgradeFilterConfig(run.Tags),
	}
}

func upgradeTemplate(use string, t *config.TemplateFileConfig) *Step {
	if t == nil {
		return nil
	}

	with := make(map[string]interface{})

	if len(t.File) > 0 {
		with["filename"] = t.File
	}

	if len(t.Template) > 0 {
		with["template"] = t.Template
	}

	if t.Variables != nil {
		with["variables"] = t.Variables
	}

	var condition string
	if t.Enabled != nil {
		condition = upgradeCondition(t)
	}

	return &Step{
		Use:  use,
		With: with,
		If:   condition,
	}
}

func upgradeVersion(t *config.TemplateFileConfig) *Step {
	if t.Enabled != nil && *t.Enabled {
		return upgradeTemplate("winchci/generate-version", t)
	} else {
		return nil
	}
}

func upgradeChangelog(t *config.TemplateFileConfig) *Step {
	return upgradeTemplate("winchci/generate-changelog", t)
}

func upgradeAsset(a *config.AssetConfig) *Step {
	if a == nil {
		return nil
	}

	with := make(map[string]interface{})
	if len(a.Filename) > 0 {
		with["filename"] = a.Filename
	}
	if len(a.Directory) > 0 {
		with["directory"] = a.Directory
	}
	if len(a.Package) > 0 {
		with["package"] = a.Package
	}
	if len(a.Variable) > 0 {
		with["variable"] = a.Variable
	}
	if len(a.Tag) > 0 {
		with["tag"] = a.Tag
	}
	if len(a.Only) > 0 {
		with["only"] = a.Only
	}
	if len(a.Except) > 0 {
		with["except"] = a.Except
	}

	var condition string
	if a.Enabled != nil {
		condition = upgradeCondition(a)
	}

	return &Step{
		Use:      "winchci/vfsgen",
		With:     with,
		If:       condition,
		Branches: upgradeFilterConfig(a.Branches),
		Tags:     upgradeFilterConfig(a.Tags),
	}
}

func upgradeTransom(a *config.TransomConfig) *Step {
	if a == nil {
		return nil
	}

	with := make(map[string]interface{})

	if len(a.Server) > 0 {
		with["server"] = a.Server
	}
	if len(a.Shipyard) > 0 {
		with["shipyard"] = a.Shipyard
	}
	if len(a.Organization) > 0 {
		with["organization"] = a.Organization
	}
	if len(a.Application) > 0 {
		with["application"] = a.Application
	}
	if len(a.Token) > 0 {
		with["token"] = a.Token
	}
	if len(a.Username) > 0 {
		with["username"] = a.Username
	}
	if len(a.Password) > 0 {
		with["password"] = a.Password
	}
	if len(a.Directory) > 0 {
		with["directory"] = a.Directory
	}
	if len(a.Artifacts) > 0 {
		with["artifacts"] = a.Artifacts
	}

	var condition string
	if a.Enabled != nil {
		condition = upgradeCondition(a)
	}

	return &Step{
		Use:      "winchci/publish-transom",
		With:     with,
		If:       condition,
		Branches: upgradeFilterConfig(a.Branches),
		Tags:     upgradeFilterConfig(a.Tags),
	}
}

func upgradeRelease(r *config.ReleaseConfig) (*Step, error) {
	if r == nil {
		return nil, nil
	}

	if len(r.Command) > 0 {
		if len(r.Input) > 0 {
			return nil, fmt.Errorf("upgrade: cannot upgrade input to commands")
		}

		return &Step{
			Run:         r.Command,
			Environment: r.Environment,
			If:          upgradeCondition(r),
			Branches:    upgradeFilterConfig(r.Branches),
			Tags:        upgradeFilterConfig(r.Tags),
		}, nil
	}

	with := make(map[string]interface{})
	if len(r.Name) > 0 {
		with["name"] = r.Name
	}
	if len(r.Artifacts) > 0 {
		with["artifacts"] = r.Artifacts
	}

	return &Step{
		Use:         "winchci/create-github-release",
		With:        with,
		Environment: r.Environment,
		Branches:    upgradeFilterConfig(r.Branches),
		Tags:        upgradeFilterConfig(r.Tags),
	}, nil
}

func upgradeHomebrew(a *config.HomebrewConfig) *Step {
	if a == nil {
		return nil
	}

	with := make(map[string]interface{})
	if len(a.Organization) > 0 {
		with["organization"] = a.Organization
	}
	if len(a.Repository) > 0 {
		with["repository"] = a.Repository
	}
	if len(a.Template) > 0 {
		with["template"] = a.Template
	}
	if len(a.File) > 0 {
		with["file"] = a.File
	}
	if len(a.Url) > 0 {
		with["url"] = a.Url
	}
	if len(a.Asset) > 0 {
		with["asset"] = a.Asset
	}
	if len(a.Install) > 0 {
		with["install"] = a.Install
	}
	if len(a.Test) > 0 {
		with["test"] = a.Test
	}
	if len(a.DependsOn) > 0 {
		with["depends-on"] = a.DependsOn
	}
	if len(a.Variables) > 0 {
		with["variables"] = a.Variables
	}

	var condition string
	if a.Enabled != nil {
		condition = upgradeCondition(a)
	}

	return &Step{
		Use:      "winchci/generate-homebrew",
		With:     with,
		If:       condition,
		Branches: upgradeFilterConfig(a.Branches),
		Tags:     upgradeFilterConfig(a.Tags),
	}
}

func upgradeDockerfile(t *config.TemplateFileConfig) *Step {
	return upgradeTemplate("winchci/generate-dockerfile", t)
}

func upgradePublishDocker(a *config.DockerConfig) *Step {
	if a == nil {
		return nil
	}

	with := make(map[string]interface{})
	if len(a.Server) > 0 {
		with["server"] = a.Server
	}
	if len(a.Organization) > 0 {
		with["organization"] = a.Organization
	}
	if len(a.Repository) > 0 {
		with["repository"] = a.Repository
	}
	if len(a.Username) > 0 {
		with["username"] = a.Username
	}
	if len(a.Password) > 0 {
		with["password"] = a.Password
	}
	if len(a.Dockerfile) > 0 {
		with["dockerfile"] = a.Dockerfile
	}
	if len(a.Context) > 0 {
		with["context"] = a.Context
	}
	if a.BuildArgs != nil {
		with["buildargs"] = a.BuildArgs
	}

	var condition string
	if a.Enabled != nil {
		condition = upgradeCondition(a)
	}

	return &Step{
		Use:      "winchci/publish-docker",
		With:     with,
		If:       condition,
		Branches: upgradeFilterConfig(a.Branches),
		Tags:     upgradeFilterConfig(a.Tags),
	}
}

func hasDefaultRunConfig(r *config.RunConfig, defaultConfig *config.RunConfig) bool {
	if r.Enabled != defaultConfig.Enabled {
		return false
	}

	if r.Name != defaultConfig.Name {
		return false
	}

	if r.Command != defaultConfig.Command {
		return false
	}

	return true

}

func removeRunConfig(r *config.RunConfig, defaultConfig *config.RunConfig) *config.RunConfig {
	if !hasDefaultRunConfig(r, defaultConfig) {
		return r
	}

	return nil
}

func hasDefaultTemplateConfig(r *config.TemplateFileConfig, defaultConfig *config.TemplateFileConfig) bool {
	if r.Enabled != defaultConfig.Enabled {
		return false
	}

	if r.File != defaultConfig.File {
		return false
	}

	if r.Template != defaultConfig.Template {
		return false
	}

	return true
}

func removeTemplateConfig(r *config.TemplateFileConfig, defaultConfig *config.TemplateFileConfig) *config.TemplateFileConfig {
	if !hasDefaultTemplateConfig(r, defaultConfig) {
		return r
	}

	return nil
}

func removeDefaultCommands(cfg *config.Config, defaultConfig *config.Config) bool {
	changed := false

	if hasDefaultRunConfig(cfg.Install, defaultConfig.Install) && hasDefaultRunConfig(cfg.Build, defaultConfig.Build) &&
		hasDefaultRunConfig(cfg.Test, defaultConfig.Test) {
		cfg.Install = removeRunConfig(cfg.Install, defaultConfig.Install)
		cfg.Build = removeRunConfig(cfg.Build, defaultConfig.Build)
		cfg.Test = removeRunConfig(cfg.Test, defaultConfig.Test)
		changed = true
	}

	if hasDefaultTemplateConfig(cfg.Version, defaultConfig.Version) {
		cfg.Version = removeTemplateConfig(cfg.Version, defaultConfig.Version)
	}

	if hasDefaultTemplateConfig(cfg.Dockerfile, defaultConfig.Dockerfile) {
		cfg.Dockerfile = removeTemplateConfig(cfg.Dockerfile, defaultConfig.Dockerfile)
	}

	return changed
}

func Upgrade(cfg *config.Config) (*Project, error) {
	cfg.Changelog = removeTemplateConfig(cfg.Changelog, &config.TemplateFileConfig{
		Template: "!changelog.tmpl",
		File:     "CHANGELOG.md",
	})

	cfg2 := &Project{
		Version:     2,
		Name:        cfg.Name,
		Description: cfg.Description,
		Repository:  cfg.Repository,
		Environment: cfg.Environment,
	}

	var build *Step

	language, toolchain := config.GuessLanguageAndToolchain()
	if cfg.Language == language && cfg.Toolchain == toolchain {
		var changed bool

		// Remove default commands
		switch language {
		case "docker":
			changed = removeDefaultCommands(cfg, config.DefaultDockerConfig)

		case "java":
			changed = removeDefaultCommands(cfg, config.DefaultJavaMavenConfig)

		case "node":
			if toolchain == "yarn" {
				changed = removeDefaultCommands(cfg, config.DefaultYarnConfig)
			} else {
				changed = removeDefaultCommands(cfg, config.DefaultNpmConfig)
			}

		case "python":
			changed = removeDefaultCommands(cfg, config.DefaultPythonConfig)

		case "scala":
			changed = removeDefaultCommands(cfg, config.DefaultScalaSbtConfig)

		default:
			changed = removeDefaultCommands(cfg, config.DefaultGoConfig)
		}

		if changed {
			build = &Step{
				Use: "winchci/" + language,
			}
			if cfg.Build != nil {
				build.Branches = upgradeFilterConfig(cfg.Build.Branches)
				build.Tags = upgradeFilterConfig(cfg.Build.Tags)
			}
		}
	}

	if cfg.Version != nil {
		v := upgradeVersion(cfg.Version)
		if v != nil {
			cfg2.Steps = append(cfg2.Steps, v)
		}
	} else {
		cfg2.Steps = append(cfg2.Steps, &Step{
			Use: "winchci/generate-version",
		})
	}

	if cfg.Changelog != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeChangelog(cfg.Changelog))
	}

	for _, a := range cfg.Assets {
		cfg2.Steps = append(cfg2.Steps, upgradeAsset(a))
	}

	if build != nil {
		cfg2.Steps = append(cfg2.Steps, build)
	}

	if cfg.BeforeInstall != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.BeforeInstall))
	}

	if cfg.Install.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.Install))
	}

	if cfg.AfterInstall.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.AfterInstall))
	}

	if cfg.BeforeBuild != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.BeforeBuild))
	}

	if cfg.Build.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.Build))
	}

	if cfg.AfterBuild.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.AfterBuild))
	}

	if cfg.BeforeTest != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.BeforeTest))
	}

	if cfg.Test.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.Test))
	}

	if cfg.AfterTest.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.AfterTest))
	}

	if cfg.BeforeRelease != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.BeforeRelease))
	}

	if cfg.Release != nil && (cfg.Release.Enabled == nil || (cfg.Release.Enabled != nil && *cfg.Release.Enabled)) {
		if cfg.Release.IsEnabled() {
			cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.Release.RunConfig()))
		}

		s, err := upgradeRelease(cfg.Release)
		if err != nil {
			return nil, err
		}

		cfg2.Steps = append(cfg2.Steps, s)
	}

	if cfg.AfterRelease.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.AfterRelease))
	}

	if cfg.BeforePublish != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.BeforePublish))
	}

	if cfg.Publish.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.Publish))
	}

	if cfg.Transom != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeTransom(cfg.Transom))
	}

	if cfg.Homebrew.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeHomebrew(cfg.Homebrew))
	}

	if cfg.Dockerfile != nil {
		cfg2.Steps = append(cfg2.Steps, upgradeDockerfile(cfg.Dockerfile))
	}

	for _, df := range cfg.Dockerfiles {
		cfg2.Steps = append(cfg2.Steps, upgradeDockerfile(df))
	}

	if cfg.Docker != nil {
		cfg2.Steps = append(cfg2.Steps, upgradePublishDocker(cfg.Docker))
	}

	for _, d := range cfg.Dockers {
		cfg2.Steps = append(cfg2.Steps, upgradePublishDocker(d))
	}

	if cfg.AfterPublish.IsEnabled() {
		cfg2.Steps = append(cfg2.Steps, upgradeRunConfig(cfg.AfterPublish))
	}

	return cfg2, nil
}
