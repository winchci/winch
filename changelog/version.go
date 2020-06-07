package changelog

import (
	"context"
	"encoding/json"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type VersionBumpInfo struct {
	Name        string
	Version     string
	Description string
	ReleaseName string
	Prerelease  string
}

func writeVersionFromTemplate(cfg *config.Config, version VersionBumpInfo) error {
	f, err := os.Create(filepath.Join(cfg.BasePath, cfg.Version.File))
	if err != nil {
		return err
	}

	defer f.Close()

	vars := cfg.Version.Variables
	if vars == nil {
		vars = make(map[string]string)
	}
	if _, ok := vars["Name"]; !ok {
		vars["Name"] = version.Name
	}
	if _, ok := vars["Version"]; !ok {
		vars["Version"] = version.Version
	}
	if _, ok := vars["Description"]; !ok {
		vars["Description"] = version.Description
	}
	if _, ok := vars["ReleaseName"]; !ok {
		vars["ReleaseName"] = version.ReleaseName
	}
	if _, ok := vars["Prerelease"]; !ok {
		vars["Prerelease"] = version.Prerelease
	}

	err = templates.Load(cfg.BasePath, cfg.Version.Template).Execute(f, vars)
	if err != nil {
		return err
	}

	return nil
}

func writeVersionNode(cfg *config.Config, version VersionBumpInfo) error {
	b, err := ioutil.ReadFile(cfg.Version.File)
	if err != nil {
		return err
	}

	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	if len(version.Description) > 0 {
		j["description"] = version.Description
	}

	j["version"] = strings.TrimPrefix(version.Version, "v")

	b, err = json.MarshalIndent(j, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cfg.Version.File, b, 0644)
}

func GetVersionFromReleases(releases []*winch.Release) (string, string) {
	var version string
	var prerelease string
	if len(releases) > 0 {
		version = releases[0].Version
		prerelease = ""
	} else {
		version = "v0.0.0"
		prerelease = "dev"
	}

	if version[0] == 'v' {
		version = version[1:]
	}

	return version, prerelease
}

func WriteVersion(cfg *config.Config, version, prerelease string) error {
	vbi := VersionBumpInfo{
		Name:        cfg.Name,
		Description: cfg.Description,
		Version:     version,
		ReleaseName: winch.Name(context.Background(), "adjectives", "animals"),
		Prerelease:  prerelease,
	}

	if cfg.Language == "node" {
		return writeVersionNode(cfg, vbi)
	} else {
		return writeVersionFromTemplate(cfg, vbi)
	}
}
