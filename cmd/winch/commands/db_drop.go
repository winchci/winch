package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"os"
	"os/exec"
	"strconv"
)

func dropdb(ctx context.Context) error {
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
		fmt.Println("Dropping database", cfg.Database.Database)
	}

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("dropdb",
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		"--if-exists",
		"-i",
		cfg.Database.Database,
	)

	c.Stdin = os.Stdin

	if cfg.Verbose {
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout
	}

	return c.Run()
}

func init() {
	var cmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop the database",
		Run:   Runner(dropdb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())

	dbCmd.AddCommand(cmd)
}
