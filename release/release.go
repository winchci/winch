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

package release

import (
	"context"
	"fmt"
	"github.com/winchci/winch"
	"github.com/winchci/winch/changelog"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
	"mime"
)

func DoRelease(ctx context.Context, cfg *config.Config) error {
	if !winch.CheckFilters(ctx, cfg.Release.Branches, cfg.Release.Tags) {
		return nil
	}

	err := mime.AddExtensionType(".gz", "application/gzip")
	if err != nil {
		return err
	}

	err = winch.Run(ctx, cfg.BeforeRelease, cfg)
	if err != nil {
		return err
	}

	// Run a provided release command
	if cfg.Release != nil {
		err = winch.Run(ctx, cfg.Release.RunConfig(), cfg)
		if err != nil {
			return err
		}
	}

	releases, err := changelog.MakeReleases(ctx, cfg)
	if err != nil {
		return err
	}

	changelog, err := winch.MakeChangelog(ctx, cfg.Repository, releases)
	if err != nil {
		return err
	}

	changelog.Releases = changelog.Releases[0:1]

	body, err := templates.Execute(ctx, "release.tmpl", changelog)
	if err != nil {
		return err
	}

	tag := changelog.Releases[0].Version

	client, err := winch.NewGitHub(ctx, cfg.Repository)
	if err != nil {
		return err
	}

	fmt.Println("Creating release")
	rel, err := client.CreateRelease(ctx, tag, body)
	if err != nil {
		return err
	}

	fmt.Println("Uploading assets")
	for _, artifact := range cfg.Release.Artifacts {
		err = client.UploadAsset(ctx, rel.GetID(), artifact)
		if err != nil {
			return err
		}
	}

	err = winch.Run(ctx, cfg.AfterRelease, cfg)
	if err != nil {
		return err
	}

	return nil
}
