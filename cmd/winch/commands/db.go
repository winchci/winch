package commands

import "github.com/spf13/cobra"

var dbCmd = &cobra.Command{
	Use:              "db",
	Short:            "Database commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
