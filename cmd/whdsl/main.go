package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bufbuild/connect-grpcreflect-go"

	"github.com/maxihafer/whdsl/cmd/whdsl/internal"
	"github.com/maxihafer/whdsl/pkg/article"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	conf := internal.NewInitializedMariaDBConfigFromEnv()
	dsn := conf.DSN()

	logrus.WithField("dsn",dsn).Info("connecting to database")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	db.AutoMigrate(&article.Article{})
	db.AutoMigrate(&article.Transaction{})

	articleRepo := article.NewArticleRepository(db)
	transactionRepo := article.NewTransactionRepository(db)

	articleService := article.NewArticleService(articleRepo)
	transactionService := article.NewTransactionService(transactionRepo)

	mux := http.NewServeMux()
	
	reflector := grpcreflect.NewStaticReflector(
		articlev1connect.ArticleServiceName,
		articlev1connect.TransactionServiceName,
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle(articlev1connect.NewArticleServiceHandler(articleService))
	mux.Handle(articlev1connect.NewTransactionServiceHandler(transactionService))

	if err := http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		return err
	}

	return nil
}
