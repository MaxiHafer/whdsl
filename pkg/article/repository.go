package article

import (
	"context"

	"whdsl/pkg/mariadb"
)

type IRepository interface{}

type Repository struct {
	db mariadb.Backend
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Article, error) {
	article := new(Article)

	err := r.db.BindByID(ctx, id, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	return r.db.Delete(ctx, &Article{ID: id})
}

func (r *Repository) Store(ctx context.Context, article *Article) error {
	return r.db.InsertOrUpdate(ctx, article)
}