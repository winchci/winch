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
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
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
