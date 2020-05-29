package commands

import (
	"context"
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"path"
	"strings"
)

func dropVaultPath(ctx context.Context, client *vault.Client, root, searchPath string) error {
	sec, err := client.Logical().List(searchPath)
	if err != nil {
		return err
	}

	if sec != nil && sec.Data != nil && sec.Data["keys"] != nil {
		for _, key := range sec.Data["keys"].([]interface{}) {
			keyPath := key.(string)
			if strings.HasSuffix(keyPath, "/") {
				err = dropVaultPath(ctx, client, root, path.Join(searchPath, keyPath))
				if err != nil {
					return err
				}
			} else {
				_, err = client.Logical().Delete(path.Join(searchPath, keyPath))
				if err != nil {
					return err
				}
			}
		}
	}

	return err
}

func dropVault(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	client, err := vault.NewClient(&vault.Config{
		Address: cfg.Vault.Address,
	})
	if err != nil {
		return err
	}

	client.SetToken(cfg.Vault.Token)

	if len(cfg.Vault.Dir) == 0 {
		cfg.Vault.Dir = "./data/vault"
	}

	err = dropVaultPath(ctx, client, cfg.Vault.Dir, cfg.Vault.Prefix)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop the database",
		Run:   Runner(dropVault),
		Args:  cobra.NoArgs,
	}

	vaultCmd.AddCommand(cmd)
}
