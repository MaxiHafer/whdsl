package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"whdsl/cmd/whdsl/internal"
	"whdsl/pkg/article"
	"whdsl/pkg/transaction"

	"whdsl/pkg/mariadb"
)

var _ internal.Server = &Server{}

type Server struct {
	srv *http.Server
}

func (s *Server) Run(ctx context.Context) error {

	models := []mariadb.Model{
		new(article.Model),
		new(transaction.Model),
	}

	mariaBackend, err := mariadb.NewInitializedBackendFromEnv(ctx,models...)
	if err != nil {
		return err
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")

	transactionRepo := transaction.NewRepository(mariaBackend)
	transactionHandler := transaction.NewHandler(v1.Group("/transactions"), transactionRepo)

	articleRepo := article.NewRepository(mariaBackend)
	articleHandler := article.NewHandler(v1.Group("/articles"), articleRepo)

	mariaBackend.ResetModel(ctx, new(article.Model))
	mariaBackend.ResetModel(ctx, new(transaction.Model))

	transactionHandler.RegisterRoutes()
	articleHandler.RegisterRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.srv = &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

	return nil
}
