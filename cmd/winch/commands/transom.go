package commands

import "github.com/spf13/cobra"

var transomCmd = &cobra.Command{
	Use:              "transom",
	Short:            "Transom commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(transomCmd)
}
