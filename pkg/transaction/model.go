package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"

	"whdsl/pkg/mariadb"
)

var _ mariadb.Model = &Transaction{}

func NewModel() *Transaction {
	return &Transaction{
		ID: uuid.NewString(),
	}
}

type Transaction struct {
	bun.BaseModel
	ID        string    `bun:",pk,type:varchar(36)"`
	ArticleID string    `bun:",notnull,type:varchar(36)"`
	Direction Direction `bun:",notnull,type:tinyint(1)"`
	Amount    int

	CreatedAt time.Time `bun:",nullzero"`
	UpdatedAt time.Time `bun:",nullzero"`
}

func (t *Transaction) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		t.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		t.UpdatedAt = time.Now()
	}
	return nil
}

func (t *Transaction) Init(ctx context.Context, db *bun.DB) error {
	return db.ResetModel(ctx, (*Transaction)(nil))
}
