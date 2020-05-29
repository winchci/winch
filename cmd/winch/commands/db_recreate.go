package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
)

func recreatedb(ctx context.Context) error {
	var err error

	err = dropdb(ctx)
	if err != nil {
		return err
	}

	err = createdb(ctx)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "recreate",
		Short: "Recreate the database",
		Run:   Runner(recreatedb),
		Args:  cobra.NoArgs,
	}

	config.AddDB(cmd.Flags())
	cmd.Flags().String("dir", "./data", "output directory")
	cmd.Flags().Bool("timestamp", false, "create a timestamped database")
	cmd.Flags().Bool("update", false, "update the connection information in application configuration")

	dbCmd.AddCommand(cmd)
}
