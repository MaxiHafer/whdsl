package api

import (
	"github.com/gin-gonic/gin"

	"whdsl/pkg/article"
)

func Aricles(g *gin.RouterGroup) {
	g.GET("/", article.List)
	g.GET("/:article_id", article.On)
}
