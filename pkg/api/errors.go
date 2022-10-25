package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrNotFound(c *gin.Context, format string, args... any) {
	c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf(format, args))
}

func ErrInternal(c *gin.Context, format string, args ... any) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf(format, args))
}

func ErrBadRequest(c *gin.Context, format string, args... any) {
	c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf(format, args))
}