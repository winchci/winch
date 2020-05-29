package winch

import (
	"context"
	"github.com/switch-bit/winch/wordlists"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
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
