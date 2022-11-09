package article

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/maxihafer/whdsl/pkg/grpc"
)

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

type TransactionRepository struct {
	db *gorm.DB
}

func (r *TransactionRepository) AssertArticleForIDPresent(id string) (bool, error) {
	agg := &Article{ID: id}
	res := r.db.First(agg)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, grpc.ErrInternal(res.Error)
		}
	}

	return true, nil
}

func (r *TransactionRepository) Store(agg *Transaction) error {
	if res := r.db.Create(agg); res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *TransactionRepository) Get(id string) (*Transaction, error) {
	filter := &Transaction{ID: id}
	if res := r.db.First(filter); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}

		return nil, grpc.ErrInternal(res.Error)
	}

	return filter, nil
}

func (r *TransactionRepository) List() ([]*Transaction, error) {
	var aggs []*Transaction
	if res := r.db.Find(&aggs); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}
	}
	
	return aggs, nil
}

func (r *TransactionRepository) Delete(id string) error {
	filter := &Transaction{ID: id}
	if res := r.db.Delete(filter); res.Error != nil {
		return grpc.ErrInternal(res.Error)
	}

	return nil
}
