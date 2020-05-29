/*
winch - Universal Build and Release Tool
Copyright (C) 2020 Switchbit, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

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
	"github.com/winchci/winch/config"
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
