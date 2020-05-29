package commands

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"os"
)

func dropDynamoDB(ctx context.Context) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	var endpoint *string
	if cfg.Dynamodb != nil {
		endpoint = aws.String(cfg.Dynamodb.GetEndpoint())
	}
	var region *string
	if r := os.Getenv("AWS_DEFAULT_REGION"); len(r) != 0 {
		region = aws.String(r)
	} else {
		region = aws.String("local")
	}

	s, err := session.NewSession(&aws.Config{
		Endpoint: endpoint,
		Region:   region,
	})
	if err != nil {
		return err
	}

	s, err = session.NewSession(&aws.Config{
		Endpoint: endpoint,
		Region:   region,
		Credentials: credentials.NewChainCredentials([]credentials.Provider{
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{},
			&ec2rolecreds.EC2RoleProvider{
				Client: ec2metadata.New(s),
			},
		}),
	})
	if err != nil {
		return err
	}

	db := dynamodb.New(s)

	o, err := db.ListTablesWithContext(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		return err
	}

	for _, tableName := range o.TableNames {
		fmt.Println("Dropping", *tableName)
		_, err := db.DescribeTable(&dynamodb.DescribeTableInput{
			TableName: tableName,
		})
		if err == nil {
			_, err := db.DeleteTable(&dynamodb.DeleteTableInput{
				TableName: tableName,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "drop",
		Short: "Drop the database",
		Run:   Runner(dropDynamoDB),
		Args:  cobra.NoArgs,
	}

	config.AddDynamodb(cmd.Flags(), "dynamodb")

	dynamodbCmd.AddCommand(cmd)
}
