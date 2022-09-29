package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "resource not found")
}

func ErrInternal(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, "internal server error")
}