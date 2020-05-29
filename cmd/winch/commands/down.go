package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func down(ctx context.Context) error {
	return fmt.Errorf("not implemented (yet)")
}

func init() {
	var cmd = &cobra.Command{
		Use:   "down",
		Short: "Take down an environment",
		Run:   Runner(down),
		Args:  cobra.NoArgs,
	}

	rootCmd.AddCommand(cmd)
}
