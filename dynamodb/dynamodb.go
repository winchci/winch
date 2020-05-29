package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

type IndexDef struct {
	Name       string
	Projection string
	Key        map[string]string
}

type TableDef struct {
	Name       string
	Key        map[string]string
	Attributes map[string]string
	Indexes    []*IndexDef
}

type DatabaseDef struct {
	Tables []*TableDef
}

func (i IndexDef) GetIndex() *dynamodb.GlobalSecondaryIndex {
	projection := &dynamodb.Projection{}

	switch i.Projection {
	case "", "ALL":
		projection.ProjectionType = aws.String("ALL")

	case "KEYS_ONLY":
		projection.ProjectionType = aws.String("KEYS_ONLY")

	default:
		projection.ProjectionType = aws.String("INCLUDE")

		for _, p := range strings.Split(i.Projection, ",") {
			projection.NonKeyAttributes = append(projection.NonKeyAttributes, aws.String(strings.TrimSpace(p)))
		}
	}

	return &dynamodb.GlobalSecondaryIndex{
		IndexName:  aws.String(i.Name),
		KeySchema:  getKeys(i.Key),
		Projection: projection,
	}
}

func (i IndexDef) GetKeys() []*dynamodb.KeySchemaElement {
	return getKeys(i.Key)
}

func (t TableDef) GetIndexes() (indexes []*dynamodb.GlobalSecondaryIndex) {
	for _, i := range t.Indexes {
		indexes = append(indexes, i.GetIndex())
	}
	return
}

func (t TableDef) GetKeys() []*dynamodb.KeySchemaElement {
	return getKeys(t.Key)
}

func (t TableDef) GetAttributes() []*dynamodb.AttributeDefinition {
	attrs := make([]*dynamodb.AttributeDefinition, 0)
	for k, v := range t.Attributes {
		attrs = append(attrs, &dynamodb.AttributeDefinition{
			AttributeName: aws.String(k),
			AttributeType: aws.String(v),
		})
	}

	return attrs
}

func getKeys(key map[string]string) []*dynamodb.KeySchemaElement {
	keys := make([]*dynamodb.KeySchemaElement, 0)

	for k, v := range key {
		if v != "HASH" {
			continue
		}
		keys = append(keys, &dynamodb.KeySchemaElement{
			AttributeName: aws.String(k),
			KeyType:       aws.String(v),
		})
	}

	for k, v := range key {
		if v == "HASH" {
			continue
		}
		keys = append(keys, &dynamodb.KeySchemaElement{
			AttributeName: aws.String(k),
			KeyType:       aws.String(v),
		})
	}

	return keys
}
