package article

import (
	"context"
	"time"

	"whdsl/pkg/transaction"

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
	ID            string `bun:",pk,type:varchar(36)" json:"id" example:"5b4d078d-14f2-4876-9f45-0ac244874d99"`
	Name          string `json:"name" example:"kebab"`
	MinimumAmount int                  `json:"min_amount" example:"1"`
	Transactions  []*transaction.Model `bun:"rel:has-many,join:id=article_id" swaggerignore:"true"`
	CreatedAt     time.Time            `bun:",nullzero" json:"created_at" example:"1985-04-12T23:20:50.52Z"`
	UpdatedAt     time.Time                  `bun:",nullzero" json:"updated_at" example:"1985-04-12T23:20:50.52Z"`
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
