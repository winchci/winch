package commands

import "github.com/spf13/cobra"

var dockerCmd = &cobra.Command{
	Use:              "docker",
	Short:            "Docker commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
