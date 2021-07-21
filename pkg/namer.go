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
	"io/ioutil"
	"log"
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

	root, err := wordlists.Assets.Open("/")
	if err != nil {
		panic(err)
	}

	files, err := root.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, asset := range files {
		f, err := wordlists.Assets.Open(asset.Name())
		if err != nil {
			log.Fatal(err)
		}

		s, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		WordLists[strings.TrimSuffix(asset.Name(), ".txt")] = strings.Split(string(s), "\n")
	}
}
