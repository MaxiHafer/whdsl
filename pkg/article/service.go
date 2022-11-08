package article

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/gin-gonic/gin"
	v1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)
var _ articlev1connect.ArticleServiceHandler = &Service{}

func NewHandlerForService() (string, gin.HandlerFunc ) {
	path, handler := articlev1connect.NewArticleServiceHandler(&Service{})
	return path, gin.WrapH(handler)
}

type Service struct {
	
}

func (s Service) GetArticle(ctx context.Context, c *connect_go.Request[v1.GetArticleRequest]) (*connect_go.Response[v1.GetArticleResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListArticles(ctx context.Context, c *connect_go.Request[v1.ListArticlesRequest]) (*connect_go.Response[v1.ListArticlesResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) CreateArticle(ctx context.Context, c *connect_go.Request[v1.CreateArticleRequest]) (*connect_go.Response[v1.CreateArticleResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateArticle(ctx context.Context, c *connect_go.Request[v1.UpdateArticleRequest]) (*connect_go.Response[v1.UpdateArticleResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteArticle(ctx context.Context, c *connect_go.Request[v1.DeleteArticleRequest]) (*connect_go.Response[v1.DeleteArticleResponse], error) {
	//TODO implement me
	panic("implement me")
}

