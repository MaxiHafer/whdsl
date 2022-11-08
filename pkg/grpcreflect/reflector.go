package grpcreflect

import (
	connectreflection "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/gin-gonic/gin"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

var (
	reflector = connectreflection.NewStaticReflector(
		articlev1connect.ArticleServiceName,
	)
)

func ReflectorV1() (string, gin.HandlerFunc){
	path, handler := connectreflection.NewHandlerV1(reflector)

	return path, gin.WrapH(handler)
}

func ReflectorV1Alpha() (string, gin.HandlerFunc){
	path, handler := connectreflection.NewHandlerV1Alpha(reflector)

	return path, gin.WrapH(handler)
}
