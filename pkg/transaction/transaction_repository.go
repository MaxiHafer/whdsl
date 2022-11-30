package transaction

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

func (r *Repository) Store(agg *Transaction) error {
	if res := r.db.Save(agg); res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repository) Get(id string) (*Transaction, error) {
	filter := &Transaction{ID: id}
	if res := r.db.First(filter); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}

		return nil, grpc.ErrInternal(res.Error)
	}

	return filter, nil
}

func (r *Repository) List() ([]*Transaction, error) {
	var aggs []*Transaction
	if res := r.db.Find(&aggs); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}
	}

	return aggs, nil
}

func (r *Repository) Delete(id string) error {
	filter := &Transaction{ID: id}
	if res := r.db.Delete(filter); res.Error != nil {
		return grpc.ErrInternal(res.Error)
	}

	return nil
}
