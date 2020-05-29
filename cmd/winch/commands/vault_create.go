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
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func createVault(ctx context.Context) error {
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
		cfg.Vault.Dir = filepath.Join("./data/vault")
	}

	err = filepath.Walk(cfg.Vault.Dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			p := strings.TrimSuffix(strings.TrimPrefix(path, cfg.Vault.Dir), ".json")

			fmt.Println(p)

			var data map[string]interface{}
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			err = json.Unmarshal(b, &data)
			if err != nil {
				return err
			}

			_, err = client.Logical().Write(p, data)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create the database",
		Run:   Runner(createVault),
		Args:  cobra.NoArgs,
	}

	vaultCmd.AddCommand(cmd)
}
