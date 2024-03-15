package common

import "context"

type Repo interface {
	RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
