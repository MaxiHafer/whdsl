package article

import (
	"context"

	"github.com/uptrace/bun"

	"whdsl/pkg/mariadb"
)

var _ mariadb.Model = Article{}

type Article struct {
	bun.BaseModel `bun:"table:articles"`
	ID            string
	Name          string
}

func (a Article) Init(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*Article)(nil)).Exec(ctx)
	return err
}
