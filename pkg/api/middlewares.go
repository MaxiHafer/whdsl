package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, e := range c.Errors {
		err := e.Err
		if errors.Is(err, sql.ErrNoRows) {
			c.String(http.StatusNotFound, "resource not found")
		}
		logrus.WithFields(logrus.Fields{
			"handler": c.HandlerName(),
			"path":    c.FullPath(),
		}).Error(errors.WithStack(err))

		c.JSON(http.StatusInternalServerError, "")
	}
}
