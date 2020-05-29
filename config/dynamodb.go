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

package config

import (
	"strings"

	"github.com/spf13/pflag"
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
