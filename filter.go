package winch

import (
	"context"
	"github.com/switch-bit/winch/config"
	"regexp"
	"strings"
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
