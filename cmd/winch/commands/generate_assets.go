package commands

import (
	"context"
	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/vfsgen"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"github.com/switch-bit/winch/config"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func generateAsset(ctx context.Context, asset *config.AssetConfig) error {
	if !asset.IsEnabled() {
		return nil
	}

	if !winch.CheckFilters(ctx, asset.Branches, asset.Tags) {
		return nil
	}

	var d http.FileSystem
	d = http.Dir(asset.Directory)

	if len(asset.Only) != 0 {
		d = filter.Keep(d, func(path string, fi os.FileInfo) bool {
			if fi.IsDir() {
				return true
			}

			for _, p := range asset.Only {
				if ok, err := regexp.MatchString(p, path); err == nil && ok {
					return true
				}
			}

			return false
		})
	}

	if len(asset.Except) != 0 {
		d = filter.Skip(d, func(path string, fi os.FileInfo) bool {
			if fi.IsDir() {
				return false
			}

			for _, p := range asset.Except {
				if ok, err := filepath.Match(p, path); err == nil && ok {
					return true
				}
			}

			return false
		})
	}

	err := vfsgen.Generate(d, vfsgen.Options{
		Filename:     asset.Filename,
		PackageName:  asset.Package,
		BuildTags:    asset.Tag,
		VariableName: asset.Variable,
	})

	return err
}

func generateAssets(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	for _, asset := range cfg.Assets {
		err = generateAsset(ctx, asset)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "assets",
		Short: "Generate assets",
		Run:   Runner(generateAssets),
		Args:  cobra.NoArgs,
	}

	generateCmd.AddCommand(cmd)
}
