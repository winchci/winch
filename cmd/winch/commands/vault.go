package commands

import "github.com/spf13/cobra"

var vaultCmd = &cobra.Command{
	Use:              "vault",
	Short:            "Vault commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(vaultCmd)
}
