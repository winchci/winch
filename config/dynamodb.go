package config

import (
	"github.com/spf13/pflag"
	"strings"
)

type DynamoDBConfig struct {
	Endpoint string
	Dir      string
}

func (c DynamoDBConfig) GetEndpoint() string {
	return c.Endpoint
}

type DynamoDBConfigIf interface {
	GetEndpoint() string
}

func makePrefixKey(key string, prefix []string) string {
	if len(prefix) == 0 {
		return key
	}

	return strings.Join(prefix, ".") + "." + key
}

// AddDynamodb adds the dynamodb parameters
func AddDynamodb(flags *pflag.FlagSet, prefix ...string) {
	flags.String(makePrefixKey("endpoint", prefix), "", "Dynamodb endpoint")
}
