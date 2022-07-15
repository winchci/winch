/*
winch - Universal Build and Release Tool
Copyright (C) 2021 Ketch Kloud, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package pkg

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/winchci/winch/wordlists"
)

var WordLists = make(map[string][]string)

func Name(_ context.Context, segments ...string) string {
	var results []string
	for _, segment := range segments {
		if list, ok := WordLists[segment]; ok {
			results = append(results, list[rand.Intn(len(list))])
		}
	}

	return strings.Join(results, " ")
}

func init() {
	rand.Seed(time.Now().Unix())

	WordLists["adjectives"] = strings.Split(wordlists.Adjectives, "\n")
	WordLists["animals"] = strings.Split(wordlists.Animals, "\n")
}
