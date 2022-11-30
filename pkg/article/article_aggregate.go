package article

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
)

func NewArticle() *Article {
	return &Article{
		ID: uuid.NewString(),
	}
}

type Article struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	MinAmount int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}



func (a *Article) ToProto() *v1.Article {
	return &v1.Article{
		Id:        a.ID,
		Name:      a.Name,
		MinAmount: a.MinAmount,
		CreatedAt: timestamppb.New(a.CreatedAt),
		UpdatedAt: timestamppb.New(a.UpdatedAt),
	}
}
