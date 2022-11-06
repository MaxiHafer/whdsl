package main

import (
	"github.com/maxihafer/whdsl/pkg/article"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	articlev1 "go.buf.build/grpc/go/maxihafer/whdsl/article/v1"
	"google.golang.org/grpc"
	"net"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	articleService := &article.Service{}

	path, handler := articlev1connect.NewArticleServiceHandler(articleService)

	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", listenOn)
	}

	server := grpc.NewServer()
	articlev1.RegisterArticleServiceServer(server, &article.Service{})
	logrus.Infof("listening on %s",listenOn)

	if err := server.Serve(listener); err != nil {
		return errors.Wrap(err, "failed to serve gRPC server")
	}

	return nil
}
