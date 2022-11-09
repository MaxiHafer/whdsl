package article

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/pkg/errors"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

var _ articlev1connect.TransactionServiceHandler = &TransactionService{}

func NewTransactionService(r *TransactionRepository) *TransactionService {
	return &TransactionService{r: r}
}

type TransactionService struct {
	r *TransactionRepository
}

func (s *TransactionService) GetTransaction(ctx context.Context, c *connect.Request[v1.GetTransactionRequest]) (*connect.Response[v1.GetTransactionResponse], error) {
	agg, err := s.r.Get(c.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.GetTransactionResponse{Transaction: agg.ToProto()}), nil
}

func (s *TransactionService) ListTransactions(ctx context.Context, c *connect.Request[v1.ListTransactionsRequest]) (*connect.Response[v1.ListTransactionsResponse], error) {
	aggs, err := s.r.List()
	if err != nil {
		return nil, err
	}

	response := connect.NewResponse(&v1.ListTransactionsResponse{
		Transactions: make([]*v1.Transaction, len(aggs)),
	})
	for i, agg := range aggs {
		response.Msg.Transactions[i] = agg.ToProto()
	}

	return response, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, c *connect.Request[v1.CreateTransactionRequest]) (*connect.Response[v1.CreateTransactionResponse], error) {
	agg := NewTransaction()

	agg.Count = c.Msg.GetCount()
	agg.Type = c.Msg.GetType()
	agg.ArticleID = c.Msg.GetArticleId()
	agg.CreatedAt = time.Now()
	agg.UpdatedAt = time.Now()

	found, err := s.r.AssertArticleForIDPresent(c.Msg.GetArticleId())
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.Errorf("no article found for id: %s", c.Msg.GetArticleId()))
	}

	if err = s.r.Store(agg); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.CreateTransactionResponse{Id: agg.ID}), nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, c *connect.Request[v1.UpdateTransactionRequest]) (*connect.Response[v1.UpdateTransactionResponse], error) {
	agg := &Transaction{}

	agg.ID = c.Msg.GetId()
	agg.Count = c.Msg.GetCount()
	agg.Type = c.Msg.GetType()

	if err := s.r.Store(agg); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.UpdateTransactionResponse{Id: agg.ID}), nil
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, c *connect.Request[v1.DeleteTransactionRequest]) (*connect.Response[v1.DeleteTransactionResponse], error) {
	if err := s.r.Delete(c.Msg.GetId()); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteTransactionResponse{Id: c.Msg.Id}), nil
}
