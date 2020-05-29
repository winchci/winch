package commands

import (
	"context"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/templates"
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

func writeVersionGo(cfg *config.Config, version VersionBumpInfo) error {
	f, err := os.Create(filepath.Join(cfg.BasePath, cfg.Version.File))
	if err != nil {
		return err
	}

	defer f.Close()

	err = templates.Load(cfg.BasePath, cfg.Version.Template).Execute(f, version)
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

var versionWriters = map[string]func(cfg *config.Config, version VersionBumpInfo) error{
	"go":   writeVersionGo,
	"node": writeVersionNode,
}

func getVersionFromReleases(releases []*winch.Release) (string, string) {
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

func writeVersion(cfg *config.Config, version, prerelease string) error {
	return versionWriters[cfg.Language](cfg, VersionBumpInfo{
		Name:        cfg.Name,
		Description: cfg.Description,
		Version:     version,
		ReleaseName: winch.Name(context.Background(), "adjectives", "animals"),
		Prerelease:  prerelease,
	})
}

func generateVersion(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	cmd := config.CommandFromContext(ctx)

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	version, prerelease := getVersionFromReleases(releases)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	if len(output) > 0 {
		cfg.Version.File = output
	}

	return writeVersion(cfg, version, prerelease)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Generate the version file",
		Run:   Runner(generateVersion),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringP("output", "o", "", "output file")

	generateCmd.AddCommand(cmd)
}
