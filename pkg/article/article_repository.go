package article

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/maxihafer/whdsl/pkg/grpc"
)

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

type ArticleRepository struct {
	db *gorm.DB
}

func (r *ArticleRepository) GetTransactionsForArticleID(id string) ([]*Transaction, error) {
	var transactions []*Transaction
	filter := &Transaction{
		ArticleID: id,
	}

	if res := r.db.Where(filter).Find(&transactions); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, grpc.ErrNotFound(res.Error)
		}
		return nil, grpc.ErrInternal(res.Error)
	}

	return transactions, nil
}

func (r *ArticleRepository) Store(agg *Article) error {
	if res := r.db.Create(agg); res.Error != nil {
		return grpc.ErrInternal(res.Error)
	}
	return nil
}

func (r *ArticleRepository) Get(id string) (*Article, error) {
	filter := &Article{ID: id}
	
	if res := r.db.First(filter); res.Error != nil {
		return nil, grpc.ErrInternal(res.Error)
	}

	return filter, nil
}

func (r *ArticleRepository) List() ([]*Article, error) {
	var aggs []*Article
	res := r.db.Find(&aggs)
	if res.Error != nil {
		return nil, grpc.ErrInternal(res.Error)
	}
	return aggs, nil
	
}

func (r *ArticleRepository) Delete(id string) error {
	filter := &Article{ID: id}

	if res := r.db.Delete(filter); res.Error != nil {
		return grpc.ErrInternal(res.Error)
	}

	return nil
}
