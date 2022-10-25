package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"

	"github.com/maxihafer/whdsl/pkg/api"
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
		middleware.OapiRequestValidator(swagger),
	)

	r = api.RegisterHandlers(r, service)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}

	return s
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	service := api.NewService()
	s := NewServer(service, *port)

	logrus.Fatal(s.ListenAndServe())
}
