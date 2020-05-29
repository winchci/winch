package commands

import (
	"bufio"
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
	"github.com/switch-bit/winch/config"
	ddb "github.com/switch-bit/winch/dynamodb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func loadDynamoDBData(ctx context.Context, db *dynamodb.DynamoDB, tableName, dir string) error {
	filename := dir + "/" + tableName + ".json"

	if fi, err := os.Stat(filename); err != nil || fi.IsDir() {
		fmt.Println("Skipping", tableName)
		return nil
	}

	fmt.Println("Loading", tableName)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(bytes.NewBuffer(b))
	line, err := reader.ReadString('\n')
	for err == nil {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			var attrs map[string]*dynamodb.AttributeValue
			err = json.Unmarshal([]byte(line), &attrs)
			if err != nil {
				return err
			}

			_, err := db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
				Item:      attrs,
				TableName: aws.String(tableName),
			})
			if err != nil {
				return err
			}
		}

		line, err = reader.ReadString('\n')
	}

	return nil
}

func createDynamoDB(ctx context.Context) error {
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

	var dbd ddb.DatabaseDef

	dir := "./data/dynamodb"
	if cfg.Dynamodb != nil && len(cfg.Dynamodb.Dir) != 0 {
		dir = cfg.Dynamodb.Dir
	}

	b, err := ioutil.ReadFile(dir + "/tables.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &dbd)
	if err != nil {
		return err
	}

	var tableNames []string
	for _, tbldef := range dbd.Tables {
		fmt.Println("Creating", tbldef.Name)

		tableNames = append(tableNames, tbldef.Name)

		dt, err := db.DescribeTable(&dynamodb.DescribeTableInput{
			TableName: aws.String(tbldef.Name),
		})
		if err == nil {
			fmt.Println(tbldef.Name, *dt.Table.TableStatus)
			continue
		}

		indexes := tbldef.GetIndexes()
		for n := range indexes {
			indexes[n].ProvisionedThroughput = &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			}
		}

		co, err := db.CreateTableWithContext(ctx, &dynamodb.CreateTableInput{
			AttributeDefinitions: tbldef.GetAttributes(),
			KeySchema:            tbldef.GetKeys(),
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			GlobalSecondaryIndexes: indexes,
			TableName:              aws.String(tbldef.Name),
		})
		if err != nil {
			return err
		}

		fmt.Println(tbldef.Name, *co.TableDescription.TableStatus)
	}

	var wg sync.WaitGroup
	for _, tableName := range tableNames {
		wg.Add(1)
		func() {
			defer wg.Done()
			err = loadDynamoDBData(ctx, db, tableName, dir)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create the database",
		Run:   Runner(createDynamoDB),
		Args:  cobra.NoArgs,
	}

	config.AddDynamodb(cmd.Flags(), "dynamodb")
	cmd.Flags().String("dir", "./data/dynamodb", "data directory")

	dynamodbCmd.AddCommand(cmd)
}
