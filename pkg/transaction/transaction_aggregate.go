package transaction

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1"
)

func NewTransaction() *Transaction {
	return &Transaction{
		ID: uuid.NewString(),
	}
}

type Transaction struct {
	ID        string `gorm:"primaryKey"`
	ArticleID string
	Type      v1.Transaction_Type
	Count     int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

func (a *Transaction) ToProto() *v1.Transaction {
	return &v1.Transaction{
		Id:        a.ID,
		ArticleId: a.ArticleID,
		Type:      a.Type,
		Count:     a.Count,
		CreatedAt: timestamppb.New(a.CreatedAt),
		UpdatedAt: timestamppb.New(a.UpdatedAt),
	}
}
