package commands

import (
	"context"
	"encoding/json"
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func listVault(ctx context.Context, client *vault.Client, root, searchPath string) error {
	err := os.MkdirAll(path.Join(root, searchPath), os.ModePerm)
	if err != nil {
		return err
	}

	sec, err := client.Logical().List(searchPath)
	if err != nil {
		return err
	}

	if sec != nil && sec.Data != nil && sec.Data["keys"] != nil {
		for _, key := range sec.Data["keys"].([]interface{}) {
			keyPath := key.(string)
			if strings.HasSuffix(keyPath, "/") {
				err = listVault(ctx, client, root, path.Join(searchPath, keyPath))
				if err != nil {
					return err
				}
			} else {
				sec, err = client.Logical().Read(path.Join(searchPath, keyPath))
				if err != nil {
					return err
				}

				if sec != nil && sec.Data != nil {
					b, err := json.MarshalIndent(sec.Data, "", "  ")
					if err != nil {
						return err
					}

					err = ioutil.WriteFile(path.Join(root, searchPath, keyPath+".json"), b, 0644)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return err
}

func dumpVault(ctx context.Context) error {
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

	err = listVault(ctx, client, cfg.Vault.Dir, cfg.Vault.Prefix)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump the database",
		Run:   Runner(dumpVault),
		Args:  cobra.NoArgs,
	}

	vaultCmd.AddCommand(cmd)
}
