package commands

import (
	"context"
	"github.com/spf13/cobra"
)

func recreateVault(ctx context.Context) error {
	err := dropVault(ctx)
	if err != nil {
		return err
	}

	err = createVault(ctx)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "recreate",
		Short: "Recreate the database",
		Run:   Runner(recreateVault),
		Args:  cobra.NoArgs,
	}

	vaultCmd.AddCommand(cmd)
}
