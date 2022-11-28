package article

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
)

func NewTransaction() *Transaction {
	return &Transaction{
		ID: uuid.NewString(),
	}
}

type Transaction struct {
	*gorm.Model

	ID        string
	ArticleID string
	Type      v1.Transaction_Type
	Count     int32
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
