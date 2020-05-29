package commands

import (
	"context"
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
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
