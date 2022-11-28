package article

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
)

func NewArticle() *Article {
	return &Article{
		ID: uuid.NewString(),
	}
}

type Article struct {
	gorm.Model
	ID string
	Name string
	MinAmount int32
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

