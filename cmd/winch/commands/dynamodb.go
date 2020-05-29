package commands

import "github.com/spf13/cobra"

var dynamodbCmd = &cobra.Command{
	Use:              "dynamodb",
	Short:            "DynamoDB commands",
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(dynamodbCmd)
}
