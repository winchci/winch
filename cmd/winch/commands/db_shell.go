package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"os"
	"os/exec"
	"strconv"
)

func shelldb(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	err = config.LoadDBConfig(ctx, cfg)
	if err != nil {
		return err
	}

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("psql",
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		cfg.Database.Database,
	)

	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}

func init() {
	var cmd = &cobra.Command{
		Use:   "shell",
		Short: "Open a database shell",
		Run:   Runner(shelldb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())

	dbCmd.AddCommand(cmd)
}
