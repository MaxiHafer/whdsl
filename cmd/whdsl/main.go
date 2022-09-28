package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"whdsl/pkg/article"
	"whdsl/pkg/transaction"

	"whdsl/pkg/mariadb"
)

// @title        WHDSL Inventory API
// @version      1.0
// @description  This is a simple api for providing inventory management capabilities
// @contact.name API Support
// @host         localhost:8080
// @BasePath     /api/v1
func main() {
	ctx := context.Background()

	models := []mariadb.Model{
		new(article.Article),
		new(transaction.Transaction),
	}

	mariaBackend, err := mariadb.NewInitializedBackendFromEnv(ctx,models...)
	if err != nil {
		logrus.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")

	transactionRepo := transaction.NewRepository(mariaBackend)
	transactionHandler := transaction.NewHandler(v1.Group("/transactions"), transactionRepo)

	articleRepo := article.NewRepository(mariaBackend)
	articleHandler := article.NewHandler(v1.Group("/articles"), articleRepo)

	mariaBackend.ResetModel(ctx, new(article.Article))
	mariaBackend.ResetModel(ctx, new(transaction.Transaction))

	transactionHandler.RegisterRoutes()
	articleHandler.RegisterRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = router.Run(":8080")
}
