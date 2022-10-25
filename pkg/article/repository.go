package article

import (
	"context"

	"whdsl/pkg/mariadb"
)

type IRepository interface{}

func NewRepository(b *mariadb.Backend) *Repository {
	return &Repository{
		db: b,
	}
}

type Repository struct {
	db *mariadb.Backend
}

func (r *Repository) ListArticles(ctx context.Context) ([]*Model, error) {
	var articles []*Model

	err := r.db.List().Model(&articles).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return articles, err
}

func (r *Repository) GetArticleByID(ctx context.Context, id string) (*Model, error) {
	article := new(Model)

	err := r.db.BindByID(ctx, id, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	return r.db.Delete(ctx, &Model{ID: id})
}

func (r *Repository) Store(ctx context.Context, article *Model) error {
	return r.db.InsertOrUpdate(ctx, article)
}