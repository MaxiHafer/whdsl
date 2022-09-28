package mariadb

import (
	"context"

	"github.com/uptrace/bun"
)

type Model interface {
	Init(ctx context.Context, db *bun.DB) error
}
