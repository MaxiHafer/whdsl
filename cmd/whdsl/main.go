package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/maxihafer/whdsl/pkg/article"
	"github.com/maxihafer/whdsl/pkg/grpcreflect"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	r := gin.New()
	r.UseH2C = true

	r.Use(
		gin.Recovery(),
		cors.Default(),
		ginlogrus.Logger(logrus.StandardLogger()),
	)

	r.Any(grpcreflect.ReflectorV1())
	r.Any(grpcreflect.ReflectorV1Alpha())

	articlePath, articleHandler := article.NewHandlerForService()

	r.POST(articlePath, articleHandler)

	if err := r.Run(":8080"); err != nil {
		return err
	}

	return nil
}
