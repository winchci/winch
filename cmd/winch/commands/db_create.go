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
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func createdb(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = config.LoadDBConfig(ctx, cfg)
	if err != nil {
		return err
	}

	if cfg.Database.Timestamp {
		cfg.Database.Database = fmt.Sprintf("%s_%s", cfg.Database.Database, time.Now().Format("20060102150405"))
	}

	if !cfg.Quiet {
		fmt.Println("Creating database", cfg.Database.Database)
	}

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("createdb",
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		"-E", "UTF8",
		cfg.Database.Database,
	)

	c.Stdin = os.Stdin

	if !cfg.Quiet {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	err = c.Run()
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Println("Creating tables")
	}

	err = filepath.Walk(cfg.Database.Dir+"/tables", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		return winch.Psql(cfg, path, cfg.Database.Database)
	})
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Println("Loading data")
	}

	err = filepath.Walk(cfg.Database.Dir+"/data", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		return winch.Psql(cfg, path, cfg.Database.Database)
	})
	if err != nil {
		return err
	}

	if cfg.Database.UpdateConfig {
		b, err := ioutil.ReadFile(fmt.Sprintf(".%s/config.yaml", cfg.Name))
		if err != nil {
			return err
		}

		var appConfig map[string]interface{}
		err = yaml.Unmarshal(b, &appConfig)
		if err != nil {
			return err
		}

		var e map[interface{}]interface{}
		if d, ok := appConfig["db"]; ok {
			e = d.(map[interface{}]interface{})
		} else {
			e = make(map[interface{}]interface{})
			appConfig["db"] = e
		}

		e["dialect"] = cfg.Database.Dialect

		if len(cfg.Database.Host) > 0 {
			e["host"] = cfg.Database.Host
		}

		e["port"] = cfg.Database.Port

		if len(cfg.Database.Database) > 0 {
			e["database"] = cfg.Database.Database
		}

		if len(cfg.Database.Username) > 0 {
			e["username"] = cfg.Database.Username
		}

		if len(cfg.Database.Password) > 0 {
			e["password"] = cfg.Database.Password
		}

		e["secure"] = cfg.Database.Secure

		b, err = yaml.Marshal(appConfig)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fmt.Sprintf(".%s/config.yaml", cfg.Name), b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create the database",
		Run:   Runner(createdb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())
	cmd.Flags().String("database.dir", "./data", "output directory")
	cmd.Flags().Bool("database.timestamp", false, "create a timestamped database")
	cmd.Flags().Bool("update", false, "update the connection information in application configuration")

	dbCmd.AddCommand(cmd)
}
