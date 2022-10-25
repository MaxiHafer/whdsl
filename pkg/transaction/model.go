package transaction

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"whdsl/pkg/mariadb"
)

var _ mariadb.Model = &Model{}

func NewModel() *Model {
	return &Model{
		ID: uuid.NewString(),
	}
}

type Model struct {
	bun.BaseModel `swaggerignore:"true"`
	ID        string    `bun:",pk,type:varchar(36)"`
	ArticleID string    `bun:",notnull,type:varchar(36)"`
	Direction Direction `bun:",notnull,type:tinyint(1)"`
	Amount    int

	CreatedAt time.Time `bun:",nullzero"`
	UpdatedAt time.Time `bun:",nullzero"`
}

func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

func (m *Model) Init(ctx context.Context, db *bun.DB) error {
	return db.ResetModel(ctx, (*Model)(nil))
}
