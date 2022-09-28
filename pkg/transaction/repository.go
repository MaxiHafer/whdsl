package transaction

import (
	"context"

	"whdsl/pkg/mariadb"
)

func NewRepository(b *mariadb.Backend) *Repository {
	return &Repository{
		db: b,
	}
}

type Repository struct {
	db *mariadb.Backend
}

func (r *Repository) GetTransactionByID(ctx context.Context, id string) (*Transaction, error) {
	transaction := new(Transaction)

	err := r.db.BindByID(ctx,id, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *Repository) ListTransactions(ctx context.Context) ([]*Transaction, error) {
	var transactions []*Transaction

	err := r.db.List().Model(&transactions).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, err
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	return r.db.Delete(ctx, &Transaction{ID: id})
}

func (r *Repository) Store(ctx context.Context, transaction *Transaction) error {
	return r.db.InsertOrUpdate(ctx, transaction)
}
