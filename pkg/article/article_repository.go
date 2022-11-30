package article

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/maxihafer/whdsl/pkg/grpc"
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Store(agg *Article) error {
	if res := r.db.Save(agg); res.Error != nil {
		return grpc.ErrInternal(res.Error)
	}
	return nil
}

func (r *Repository) GetByName(name string) (*Article, error) {
	filter := &Article{Name: name}
	article, err := r.get(filter)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *Repository) GetByID(id string) (*Article, error) {
	filter := &Article{ID: id}
	article, err := r.get(filter)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *Repository) get(filter *Article) (*Article, error) {
	if res := r.db.Take(filter, filter); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}
		return nil, grpc.ErrInternal(res.Error)
	}

	return filter, nil
}

func (r *Repository) List() ([]*Article, error) {
	var aggs []*Article
	if res := r.db.Find(&aggs); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}
		return nil, grpc.ErrInternal(res.Error)
	}

	return aggs, nil
}

func (r *Repository) Delete(id string) error {
	filter := &Article{ID: id}

	if res := r.db.Delete(filter); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return grpc.ErrNotFound(res.Error)
		}
		return grpc.ErrInternal(res.Error)
	}

	return nil
}
