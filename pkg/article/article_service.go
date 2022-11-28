package article

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

var _ articlev1connect.ArticleServiceHandler = &ArticleService{}

func NewArticleService(r *ArticleRepository) *ArticleService {
	return &ArticleService{r: r}
}

type ArticleService struct {
	r *ArticleRepository
}

func (s *ArticleService) ShowTransactions(ctx context.Context, c *connect.Request[v1.ShowTransactionsRequest]) (*connect.Response[v1.ShowTransactionsResponse], error) {
	aggs, err := s.r.GetTransactionsForArticleID(c.Msg.GetId())
	if err != nil {
		return nil, err
	}

	transactions := make([]*v1.Transaction, len(aggs))
	for i, agg := range aggs {
		transactions[i] = agg.ToProto()
	}

	return connect.NewResponse(&v1.ShowTransactionsResponse{Transactions: transactions}), nil
}

func (s *ArticleService) CalculateAmount(ctx context.Context, c *connect.Request[v1.CalculateAmountRequest]) (*connect.Response[v1.CalculateAmountResponse], error) {
	aggs, err := s.r.GetTransactionsForArticleID(c.Msg.GetId())
	if err != nil {
		return nil, err
	}

	var amount int32
	for _, agg := range aggs {
		switch agg.Type {
		case v1.Transaction_TYPE_IN:
			amount += agg.Count
		case v1.Transaction_TYPE_OUT:
			amount -= agg.Count
		}
	}

	return connect.NewResponse(&v1.CalculateAmountResponse{Amount: amount}), nil
}

func (s *ArticleService) ShowDetails(ctx context.Context, c *connect.Request[v1.ShowDetailsRequest]) (*connect.Response[v1.ShowDetailsResponse], error) {
	agg, err := s.r.Get(c.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ShowDetailsResponse{Article: agg.ToProto()}),nil
}

func (s *ArticleService) ShowAll(ctx context.Context, c *connect.Request[v1.ShowAllRequest]) (*connect.Response[v1.ShowAllResponse], error) {
	aggs, err := s.r.List()
	if err != nil {
		return nil, err
	}
	
	articles := make([]*v1.Article, len(aggs))
	for i, agg := range aggs {
		articles[i] = agg.ToProto()
	}

	return connect.NewResponse(&v1.ShowAllResponse{Articles: articles}), nil
}

func (s *ArticleService) NewArticle(ctx context.Context, c *connect.Request[v1.NewArticleRequest]) (*connect.Response[v1.NewArticleResponse], error) {
	agg := NewArticle()

	agg.Name = c.Msg.GetName()
	agg.MinAmount = c.Msg.GetMinAmount()
	agg.CreatedAt = time.Now()
	agg.UpdatedAt = time.Now()

	err := s.r.Store(agg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.NewArticleResponse{Id: agg.ID}), nil
}

func (s *ArticleService) UpdateDetails(ctx context.Context, c *connect.Request[v1.UpdateDetailsRequest]) (*connect.Response[v1.UpdateDetailsResponse], error) {
	agg := &Article{}

	agg.ID = c.Msg.GetId()
	agg.Name = c.Msg.GetName()
	agg.MinAmount = c.Msg.GetMinAmount()

	if err := s.r.Store(agg); err != nil {
		return nil, err
	}
	
	return connect.NewResponse(&v1.UpdateDetailsResponse{Id: agg.ID}), nil
}

func (s *ArticleService) RemoveArticle(ctx context.Context, c *connect.Request[v1.RemoveArticleRequest]) (*connect.Response[v1.RemoveArticleResponse], error) {
	 if err := s.r.Delete(c.Msg.GetId()); err != nil {
		 return nil, err
	 }

	return connect.NewResponse(&v1.RemoveArticleResponse{Id: c.Msg.Id}), nil
}
