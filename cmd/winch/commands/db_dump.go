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
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	"os"
	"os/exec"
	"strconv"
)

func dumpTable(cfg *config.Config, tableName string) error {
	if !cfg.Quiet {
		fmt.Println("Dumping schema for " + tableName)
	}

	filename := fmt.Sprintf("%s/tables/%s.sql", cfg.Database.Dir, tableName)

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}
	c := exec.Command("pg_dump",
		"-d", cfg.Database.Database,
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		"-t", tableName,
		"-f", filename,
		"-b",
		"-O",
		"-s",
		"-E", "UTF8",
		"--column-inserts",
		"--no-comments",
		"--no-publications",
		"--no-security-labels",
		"--no-subscriptions",
		"--no-synchronized-snapshots",
		"--no-tablespaces",
		"--no-unlogged-table-data",
		"--quote-all-identifiers",
	)

	c.Stdin = os.Stdin

	if cfg.Verbose {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	return c.Run()
}

func dumpData(cfg *config.Config, tableName string) error {
	if !cfg.Quiet {
		fmt.Println("Dumping data for " + tableName)
	}

	filename := fmt.Sprintf("%s/data/%s.sql", cfg.Database.Dir, tableName)

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("pg_dump",
		"-d", cfg.Database.Database,
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		"-t", tableName,
		"-f", filename,
		"-b",
		"-O",
		"-a",
		"-E", "UTF8",
		"--column-inserts",
		"--no-comments",
		"--no-publications",
		"--no-security-labels",
		"--no-subscriptions",
		"--no-synchronized-snapshots",
		"--no-tablespaces",
		"--no-unlogged-table-data",
		"--quote-all-identifiers",
	)

	c.Stdin = os.Stdin

	if cfg.Verbose {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	return c.Run()
}

func dumpdb(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = config.LoadDBConfig(ctx, cfg)
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Println("Dumping database", cfg.Database)
	}

	err = os.MkdirAll(cfg.Database.Dir+"/tables", os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(cfg.Database.Dir+"/data", os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	if !cfg.Quiet {
		fmt.Println("Connecting to", cfg.Database.Database)
	}

	db, err := sql.Open(cfg.Database.Dialect, cfg.Database.String())
	if err != nil {
		return err
	}
	defer db.Close()

	if !cfg.Quiet {
		fmt.Println("Listing tables")
	}

	r, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE';")
	if err != nil {
		return err
	}
	defer r.Close()

	var tables []string

	for r.Next() {
		var tableName string
		err = r.Scan(&tableName)
		if err != nil {
			return err
		}

		tables = append(tables, tableName)
	}

	for _, tableName := range tables {
		err = dumpTable(cfg, tableName)
		if err != nil {
			return err
		}
	}

	for _, tableName := range tables {
		err = dumpData(cfg, tableName)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump the database",
		Run:   Runner(dumpdb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())
	cmd.Flags().String("database.dir", "./data", "output directory")

	dbCmd.AddCommand(cmd)
}
