package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"path"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest --config=types.cfg.yaml ../../whdsl-api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest --config=server.cfg.yaml ../../whdsl-api.yaml

var _ ServerInterface = &Service{}

func NewService() *Service {
	return &Service{}
}

type Service struct {
}

func (s *Service) GetSwagger(c *gin.Context) {
	log.Println("trying to serve openapi-spec")
	wd, err := os.Getwd()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.File(path.Join(wd,"whdsl-api.yaml"))
}

func (s *Service) GetMetrics(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}

func (s *Service) GetArticles(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) PostArticles(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) DeleteArticlesId(c *gin.Context, id string) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) GetArticlesId(c *gin.Context, id string) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) PutArticlesId(c *gin.Context, id string) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) GetTransactions(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) PostTransactions(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) DeleteTransactionsId(c *gin.Context, id string) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) GetTransactionsId(c *gin.Context, id string) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) PutTransactionsId(c *gin.Context, id string) {
	// TODO implement me
	panic("implement me")
}
