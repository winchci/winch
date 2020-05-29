package commands

import "github.com/spf13/cobra"

var hookCmd = &cobra.Command{
	Use:              "hook",
	Short:            "Git hook commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
