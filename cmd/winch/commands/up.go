package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func up(ctx context.Context) error {
	return fmt.Errorf("not implemented (yet)")
}

func init() {
	var cmd = &cobra.Command{
		Use:   "up",
		Short: "Bring up an environment",
		Run:   Runner(up),
		Args:  cobra.NoArgs,
	}

	cmd.Flags().StringSlice("except", nil, "image to exclude")
	cmd.Flags().StringSlice("only", nil, "image to spin up")

	rootCmd.AddCommand(cmd)
}
