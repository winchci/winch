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

package commands

import (
	"context"
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/transom"
	"github.com/winchci/winch/version"
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
