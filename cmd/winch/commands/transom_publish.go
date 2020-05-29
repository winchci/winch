package commands

import (
	"context"
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"github.com/switch-bit/winch/transom"
	"github.com/switch-bit/winch/version"
	"io/ioutil"
	"os"
	"path"
)

func publishToTransom(ctx context.Context, cfg *config.Config, ver string) error {
	if len(cfg.Transom.Directory) == 0 && len(cfg.Transom.Artifacts) == 0 {
		return fmt.Errorf("must specify directory or artifacts")
	}

	if !winch.CheckFilters(ctx, cfg.Transom.Branches, cfg.Transom.Tags) {
		return nil
	}

	t, err := transom.NewTransom(cfg.Transom, cfg.Name)
	if err != nil {
		return err
	}

	if len(cfg.Transom.Token) > 0 {
		t.SetToken(cfg.Transom.Token)
	} else {
		if !cfg.Quiet {
			fmt.Println("Logging into Transom")
		}

		resp, err := t.Login(ctx, &transom.LoginRequest{
			Username:   cfg.Transom.Username,
			Password:   cfg.Transom.Password,
			ClientCode: "winch",
		})
		if err != nil {
			return err
		}

		t.SetToken(resp.Token)
	}

	artifactsArchive := "artifacts.zip"

	dir, err := ioutil.TempDir("", version.Name)
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	artifactsArchive = path.Join(dir, artifactsArchive)

	if !cfg.Quiet {
		fmt.Println("Creating artifacts archive")
	}

	artifacts := cfg.Transom.Artifacts
	if len(artifacts) == 0 {
		artifacts = []string{"."}
	}

	currentDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	if len(cfg.Transom.Directory) > 0 {
		err = os.Chdir(cfg.Transom.Directory)
		if err != nil {
			return err
		}
	}

	err = archiver.Archive(artifacts, artifactsArchive)
	if err != nil {
		return err
	}

	err = os.Chdir(currentDirectory)
	if err != nil {
		return err
	}

	if cfg.Verbose {
		archiver.Walk(artifactsArchive, func(f archiver.File) error {
			fmt.Println("+", f.Name())
			return nil
		})
	}

	asset, err := ioutil.ReadFile(artifactsArchive)
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Printf("Publishing version %s to Transom\n", ver)
	}
	resp, err := t.Publish(ctx, &transom.PublishRequest{
		Org:      cfg.Transom.Organization,
		App:      cfg.Transom.Application,
		Contents: asset,
		Version:  ver,
	})
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Println("Org:     ", resp.Version.Org)
		fmt.Println("App:     ", resp.Version.App)
		fmt.Println("Version: ", resp.Version.Version)
		fmt.Println("Checksum:", resp.Version.Checksum)
	}

	return nil
}

func transomPublish(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	releases, err := makeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	v, _ := getVersionFromReleases(releases)

	return publishToTransom(ctx, cfg, v)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "Publish artifacts to Transom",
		Run:   Runner(transomPublish),
		Args:  cobra.NoArgs,
	}

	transomCmd.AddCommand(cmd)
}
