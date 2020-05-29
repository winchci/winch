package winch

import "context"

type Repository interface {
	GetTags(ctx context.Context) (map[string]string, error)
	GetCommits(ctx context.Context) ([]*Commit, error)
}
