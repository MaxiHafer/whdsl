package internal

import (
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/maxihafer/whdsl/pkg/api"
	"github.com/maxihafer/whdsl/pkg/metrics"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"net/http"
	"os"
	"path"
)

func NewServer(service *api.Service, port int) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	swagger.Servers = nil

	r := gin.New()

	r.Use(
		ginlogrus.Logger(logrus.StandardLogger()),
		gin.Recovery(),
		cors.Default(),
	)

	customHandlers := r.Group("/")
	customHandlers.GET("/metrics", metrics.PrometheusHandler())

	base, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	customHandlers.StaticFile("/swagger/openapi.yaml", path.Join(base, "openapi.yaml"))

	apiHandlers := r.Group("/api/v1")
	apiHandlers.Use(middleware.OapiRequestValidator(swagger))

	apiHandlers = api.RegisterHandlers(apiHandlers,service)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}

	return s
}