package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"whdsl/pkg/mariadb"
)

func main() {


	mariaBackend, err := mariadb.NewInitializedBackendFromEnv()
	if err != nil {
		logrus.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/v1")
	{

	}
}
