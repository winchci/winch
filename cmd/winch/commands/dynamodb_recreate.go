package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
)

func recreateDynamoDB(ctx context.Context) error {
	err := dropDynamoDB(ctx)
	if err != nil {
		return err
	}

	return createDynamoDB(ctx)
}

func init() {
	var cmd = &cobra.Command{
		Use:   "recreate",
		Short: "Recreate the database",
		Run:   Runner(recreateDynamoDB),
		Args:  cobra.NoArgs,
	}

	config.AddDynamodb(cmd.Flags(), "dynamodb")
	cmd.Flags().String("dir", "./data/dynamodb", "data directory")

	dynamodbCmd.AddCommand(cmd)
}
