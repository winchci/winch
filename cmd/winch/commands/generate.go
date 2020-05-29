package commands

import "github.com/spf13/cobra"

var generateCmd = &cobra.Command{
	Use:              "generate",
	Short:            "Code generation commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
