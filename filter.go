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
	"regexp"
	"strings"

	"github.com/winchci/winch/config"
)

func getFilterRegexp(in string) (*regexp.Regexp, error) {
	in = strings.TrimSuffix(strings.TrimPrefix(in, "/"), "/")
	return regexp.Compile(in)
}

func CheckFilter(ctx context.Context, filter *config.FilterConfig, in string) bool {
	if filter == nil {
		return true
	}

	if len(filter.Ignore) > 0 {
		r, err := getFilterRegexp(filter.Ignore)
		if err != nil {
			return false
		}

		return !r.MatchString(in)
	}

	if len(filter.Only) > 0 {
		r, err := getFilterRegexp(filter.Only)
		if err != nil {
			return false
		}

		return r.MatchString(in)
	}

	return false
}

func CheckFilters(ctx context.Context, branches *config.FilterConfig, tags *config.FilterConfig) bool {
	if !tags.IsEnabled() && !branches.IsEnabled() {
		return true
	}

	ci := &CircleCI{}
	head, err := ci.GetHead(ctx)
	if err != nil {
		repository, err := NewGit(ctx, FindGitDir(ctx))
		if err != nil {
			return false
		}

		head, err = repository.GetHead(ctx)
		if err != nil {
			return false
		}
	}

	if head.IsBranch() && !CheckFilter(ctx, branches, head.GetName()) {
		return false
	}

	if head.IsTag() && !CheckFilter(ctx, tags, head.GetName()) {
		return false
	}

	return true
}
