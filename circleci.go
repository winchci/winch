package winch

import (
	"context"
	"fmt"
	"os"
)

type CircleCI struct{}

func (c CircleCI) GetHead(_ context.Context) (GitRef, error) {
	if os.Getenv("CIRCLECI") == "true" {
		tag := os.Getenv("CIRCLE_TAG")
		branch := os.Getenv("CIRCLE_BRANCH")

		if len(tag) > 0 {
			return &Tag{tag}, nil
		}

		if len(branch) > 0 {
			return &Branch{branch}, nil
		}
	}

	return nil, fmt.Errorf("not running within CircleCI")
}
