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

package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	winch "github.com/winchci/winch/pkg"
	"github.com/winchci/winch/pkg/config"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type Job struct {
	Dir      string
	Filename string
}

type builder func(ctx context.Context, cfg *config.Config, job Job, width int) error

func monoBuild(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, jobs chan Job, errC chan error, build builder, width int) {
	defer wg.Done()

	for job := range jobs {
		if err := build(ctx, cfg, job, width); err != nil {
			errC <- err
			return
		}
	}
}

func simple(ctx context.Context, cfg *config.Config, job Job, width int) error {
	c := exec.Command("winch", "ci", "--incremental", "-f", job.Filename)
	c.Dir = job.Dir
	c.Stdout = winch.NewLogTailer(os.Stdout, fmt.Sprintf("%s |", winch.ColorName(winch.PadName(job.Dir, width))))
	c.Stderr = winch.NewLogTailer(os.Stderr, fmt.Sprintf("%s |", winch.ColorName(winch.PadName(job.Dir, width))))
	return c.Run()
}

func mono(ctx context.Context, cfg *config.Config, commits []*winch.Commit) error {
	fmt.Println("Starting an incremental build on monorepo")
	fmt.Printf("Parallelism: %d\n", cfg.Parallelism)

	width := 0
	affectedPaths := make(map[string]bool)
	for _, commit := range commits {
		if len(commit.Tag) == 0 {
			for _, affectedPath := range commit.AffectedPaths {
				if len(affectedPath) > width {
					width = len(affectedPath)
				}
				affectedPaths[affectedPath] = true
			}
		}
	}

	files, err := filepath.Glob("*/winch.yml")
	if err != nil {
		return err
	}

	if cfg.Parallelism == 0 {
		cfg.Parallelism = 3
	}

	C := make(chan Job)
	errC := make(chan error, cfg.Parallelism)
	wg := sync.WaitGroup{}

	for i := 0; i < cfg.Parallelism; i++ {
		wg.Add(1)
		go monoBuild(ctx, &wg, cfg, C, errC, simple, width)
	}

	for _, file := range files {
		if affectedPaths[filepath.Dir(file)] {
			job := Job{
				Dir:      filepath.Dir(file),
				Filename: filepath.Base(file),
			}

			C <- job
		}
	}
	close(C)

	wg.Wait()

	for {
		select {
		case err := <-errC:
			return err

		default:
			return nil
		}
	}
}

func ci(ctx context.Context) error {
	if os.Getenv("CI") != "true" {
		return fmt.Errorf("must be running in a CI environment")
	}

	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	// CI is always verbose and never quiet
	cfg.Verbose = true
	cfg.Quiet = false

	var version, prerelease string

	cmd := config.CommandFromContext(ctx)

	incremental, err := cmd.Flags().GetBool("incremental")
	if err != nil {
		return err
	}

	if !incremental {
		releases, commits, err := makeReleases(ctx, cfg)
		if err != nil {
			return err
		}

		version, prerelease = getVersionFromReleases(releases)
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Prerelease: %s\n", prerelease)

		err = os.Setenv("BUILD_VERSION", version)
		if err != nil {
			return err
		}

		err = os.Setenv("BUILD_PRERELEASE", prerelease)
		if err != nil {
			return err
		}

		if cfg.Changelog.IsEnabled() {
			fmt.Println("Creating changelog")
			err = writeChangelog(ctx, cfg, releases)
			if err != nil {
				return err
			}
		}

		fmt.Println("Creating version")
		err = writeVersion(cfg, version, prerelease)
		if err != nil {
			return err
		}

		if cfg.Mono {
			err = mono(ctx, cfg, commits)
			if err != nil {
				return err
			}
		}
	} else {
		version = os.Getenv("BUILD_VERSION")
		prerelease = os.Getenv("BUILD_PRERELEASE")

		fmt.Println("Creating version")
		err = writeVersion(cfg, version, prerelease)
		if err != nil {
			return err
		}
	}

	if cfg.Install.IsEnabled() {
		fmt.Println("Installing")

		err = winch.Run(ctx, cfg.BeforeInstall, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.Install, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.AfterInstall, cfg)
		if err != nil {
			return err
		}
	}

	if len(cfg.Assets) > 0 {
		fmt.Println("Creating assets")
		for _, asset := range cfg.Assets {
			err = generateAsset(ctx, asset)
			if err != nil {
				return err
			}
		}
	}

	if cfg.Build.IsEnabled() {
		fmt.Println("Building")

		err = winch.Run(ctx, cfg.BeforeBuild, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.Build, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.AfterBuild, cfg)
		if err != nil {
			return err
		}
	}

	if cfg.Test.IsEnabled() {
		fmt.Println("Testing")

		err = winch.Run(ctx, cfg.BeforeTest, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.Test, cfg)
		if err != nil {
			return err
		}

		err = winch.Run(ctx, cfg.AfterTest, cfg)
		if err != nil {
			return err
		}
	}

	if cfg.Release.IsEnabled() && !incremental {
		fmt.Println("Releasing")
		err = release2(ctx, cfg)
		if err != nil {
			return err
		}
	}

	err = publish2(ctx, cfg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "ci",
		Short: "Execute a CI build",
		Run:   Runner(ci),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().Bool("incremental", false, "perform an incremental build")

	rootCmd.AddCommand(cmd)
}
