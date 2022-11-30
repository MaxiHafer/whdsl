package transaction

import (
	"context"

	"github.com/bufbuild/connect-go"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1"

	"github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1/transactionv1connect"
)

var _ transactionv1connect.TransactionServiceHandler = &Service{}

func NewService(r *Repository) *Service {
	return &Service{r: r}
}

type Service struct {
	r *Repository
}

func (s *Service) ShowDetails(ctx context.Context, c *connect.Request[v1.ShowDetailsRequest]) (*connect.Response[v1.ShowDetailsResponse], error) {
	agg, err := s.r.Get(c.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ShowDetailsResponse{Transaction: agg.ToProto()}), nil
}

func (s *Service) ShowAll(ctx context.Context, c *connect.Request[v1.ShowAllRequest]) (*connect.Response[v1.ShowAllResponse], error) {
	aggs, err := s.r.List()
	if err != nil {
		return nil, err
	}

	response := connect.NewResponse(
		&v1.ShowAllResponse{
			Transactions: make([]*v1.Transaction, len(aggs)),
		},
	)
	for i, agg := range aggs {
		response.Msg.Transactions[i] = agg.ToProto()
	}

	return response, nil
}

func (s *Service) NewTransaction(ctx context.Context, c *connect.Request[v1.NewTransactionRequest]) (*connect.Response[v1.NewTransactionResponse], error) {
	agg := NewTransaction()

	agg.Count = c.Msg.GetCount()
	agg.Type = c.Msg.GetType()
	agg.ArticleID = c.Msg.GetArticleId()
	
	if err := s.r.Store(agg); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.NewTransactionResponse{Id: agg.ID}), nil
}

func (s *Service) UpdateDetails(ctx context.Context, c *connect.Request[v1.UpdateDetailsRequest]) (*connect.Response[v1.UpdateDetailsResponse], error) {
	agg := &Transaction{}

	agg.ID = c.Msg.GetId()
	agg.Count = c.Msg.GetCount()
	agg.Type = c.Msg.GetType()

	if err := s.r.Store(agg); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.UpdateDetailsResponse{Id: agg.ID}), nil
}

func (s *Service) DeleteTransaction(ctx context.Context, c *connect.Request[v1.DeleteTransactionRequest]) (*connect.Response[v1.DeleteTransactionResponse], error) {
	if err := s.r.Delete(c.Msg.GetId()); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteTransactionResponse{Id: c.Msg.Id}), nil
}
