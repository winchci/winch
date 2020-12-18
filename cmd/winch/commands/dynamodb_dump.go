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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
	"github.com/winchci/winch/config"
	ddb "github.com/winchci/winch/dynamodb"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

type JsonAttributeValue struct {
	B    []byte                         `json:",omitempty"`
	BOOL *bool                          `json:",omitempty"`
	BS   [][]byte                       `json:",omitempty"`
	L    []*JsonAttributeValue          `json:",omitempty"`
	M    map[string]*JsonAttributeValue `json:",omitempty"`
	N    *string                        `json:",omitempty"`
	NS   []*string                      `json:",omitempty"`
	NULL *bool                          `json:",omitempty"`
	S    *string                        `json:",omitempty"`
	SS   []*string                      `json:",omitempty"`
}

func convertList(item []*dynamodb.AttributeValue) []*JsonAttributeValue {
	result := make([]*JsonAttributeValue, 0, len(item))
	for _, i := range item {
		result = append(result, convertItem(i))
	}
	return result
}

func convertItem(item *dynamodb.AttributeValue) *JsonAttributeValue {
	return &JsonAttributeValue{
		B:    item.B,
		BOOL: item.BOOL,
		BS:   item.BS,
		L:    convertList(item.L),
		M:    convert(item.M),
		N:    item.N,
		NS:   item.NS,
		NULL: item.NULL,
		S:    item.S,
		SS:   item.SS,
	}
}

func convert(item map[string]*dynamodb.AttributeValue) map[string]*JsonAttributeValue {
	result := make(map[string]*JsonAttributeValue)
	for k, v := range item {
		result[k] = convertItem(v)
	}
	return result
}

func dumpDynamoDBData(ctx context.Context, db *dynamodb.DynamoDB, tableName, dir string) error {
	fmt.Println("Dumping data", tableName)
	buf := bytes.NewBuffer(nil)
	e := json.NewEncoder(buf)

	err := db.ScanPagesWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}, func(so *dynamodb.ScanOutput, done bool) bool {
		for _, i := range so.Items {
			if err := e.Encode(convert(i)); err != nil {
				return false
			}
		}
		return true
	})
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dir+"/"+tableName+".json", buf.Bytes(), 0644)
}

func dumpDynamoDB(ctx context.Context) error {
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

	dbdef := &ddb.DatabaseDef{}

	for _, t := range o.TableNames {
		fmt.Println("Dumping structure", *t)

		dt, err := db.DescribeTable(&dynamodb.DescribeTableInput{
			TableName: t,
		})
		if err != nil {
			return err
		}

		tbldef := &ddb.TableDef{
			Name:       *t,
			Key:        make(map[string]string),
			Attributes: make(map[string]string),
		}

		for _, attr := range dt.Table.AttributeDefinitions {
			tbldef.Attributes[*attr.AttributeName] = *attr.AttributeType
		}

		for _, attr := range dt.Table.KeySchema {
			tbldef.Key[*attr.AttributeName] = *attr.KeyType
		}

		for _, gsi := range dt.Table.GlobalSecondaryIndexes {
			indexDef := &ddb.IndexDef{
				Name: aws.StringValue(gsi.IndexName),
				Key:  make(map[string]string),
			}

			if gsi.Projection.ProjectionType == nil {
				gsi.Projection.ProjectionType = aws.String("ALL")
			}

			switch *gsi.Projection.ProjectionType {
			case "ALL":
				indexDef.Projection = "ALL"

			case "KEYS_ONLY":
				indexDef.Projection = "KEYS_ONLY"

			case "INCLUDE":
				var columns []string
				for _, a := range gsi.Projection.NonKeyAttributes {
					columns = append(columns, aws.StringValue(a))
				}
				indexDef.Projection = strings.Join(columns, ", ")
			}

			for _, attr := range gsi.KeySchema {
				indexDef.Key[*attr.AttributeName] = *attr.KeyType
			}

			tbldef.Indexes = append(tbldef.Indexes, indexDef)
		}

		if len(tbldef.Indexes) == 0 {
			tbldef.Indexes = nil
		}

		dbdef.Tables = append(dbdef.Tables, tbldef)
	}

	b, err := yaml.Marshal(dbdef)
	if err != nil {
		return err
	}

	dir := "./data/dynamodb"
	if cfg.Dynamodb != nil && len(cfg.Dynamodb.Dir) != 0 {
		dir = cfg.Dynamodb.Dir
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dir+"/tables.yml", b, 0644)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, t := range o.TableNames {
		wg.Add(1)
		go func(tableName string) {
			defer wg.Done()
			err = dumpDynamoDBData(ctx, db, tableName, dir)
			if err != nil {
				log.Fatal(err)
			}
		}(*t)
	}
	wg.Wait()

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump the database",
		Run:   Runner(dumpDynamoDB),
		Args:  cobra.NoArgs,
	}

	config.AddDynamodb(cmd.Flags(), "dynamodb")
	cmd.Flags().String("dynamodb.dir", "./data/dynamodb", "data directory")

	dynamodbCmd.AddCommand(cmd)
}
