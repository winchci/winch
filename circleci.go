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
