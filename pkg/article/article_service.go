package article

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

var _ articlev1connect.ArticleServiceHandler = &Service{}

func NewService(r *Repository) *Service {
	return &Service{r: r}
}

type Service struct {
	r *Repository
}

func (s *Service) ShowDetailsForName(ctx context.Context, c *connect.Request[v1.ShowDetailsForNameRequest]) (*connect.Response[v1.ShowDetailsForNameResponse], error) {
	agg, err := s.r.GetByName(c.Msg.GetName())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ShowDetailsForNameResponse{Article: agg.ToProto()}), nil
}

func (s *Service) ShowDetails(ctx context.Context, c *connect.Request[v1.ShowDetailsRequest]) (*connect.Response[v1.ShowDetailsResponse], error) {
	agg, err := s.r.GetByID(c.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ShowDetailsResponse{Article: agg.ToProto()}),nil
}

func (s *Service) ShowAll(ctx context.Context, c *connect.Request[v1.ShowAllRequest]) (*connect.Response[v1.ShowAllResponse], error) {
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

func (s *Service) NewArticle(ctx context.Context, c *connect.Request[v1.NewArticleRequest]) (*connect.Response[v1.NewArticleResponse], error) {
	agg := NewArticle()

	agg.Name = c.Msg.GetName()
	agg.MinAmount = c.Msg.GetMinAmount()
	agg.CreatedAt = time.Now().UTC()
	agg.UpdatedAt = time.Now().UTC()

	if err := s.r.Store(agg); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.NewArticleResponse{Id: agg.ID}), nil
}


func (s *Service) UpdateDetails(ctx context.Context, c *connect.Request[v1.UpdateDetailsRequest]) (*connect.Response[v1.UpdateDetailsResponse], error) {
	agg := &Article{}

	agg.ID = c.Msg.GetId()
	agg.Name = c.Msg.GetName()
	agg.MinAmount = c.Msg.GetMinAmount()

	if err := s.r.Store(agg); err != nil {
		return nil, err
	}
	
	return connect.NewResponse(&v1.UpdateDetailsResponse{Id: agg.ID}), nil
}

func (s *Service) RemoveArticle(ctx context.Context, c *connect.Request[v1.RemoveArticleRequest]) (*connect.Response[v1.RemoveArticleResponse], error) {
	 if err := s.r.Delete(c.Msg.GetId()); err != nil {
		 return nil, err
	 }

	return connect.NewResponse(&v1.RemoveArticleResponse{Id: c.Msg.Id}), nil
}

func (s *Service) isNew(article *Article) (bool, error){
	article, err := s.r.GetByName(article.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	} else if errors.Is(err, gorm.ErrRecordNotFound){
		return true, nil
	} else {
		return false, nil
	}
}
