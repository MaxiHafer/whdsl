package internal

import (
	"net/http"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/maxihafer/whdsl/pkg/article"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1/transactionv1connect"
	"github.com/maxihafer/whdsl/pkg/transaction"
)

func NewServerFromEnv() (*Server, error) {
	s := &Server{}

	conf := NewMariaDBConfigFromEnv()

	s.conf = conf

	return s, nil
}

type Server struct {
	conf *MariadbConfig

	db                 *gorm.DB
	articleRepo        *article.Repository
	transactionRepo    *transaction.Repository
	articleService     *article.Service
	transactionService *transaction.Service
}

func (s *Server) runMigrations() error {

	if err := s.db.Migrator().DropTable(&article.Article{}, &transaction.Transaction{}); err != nil {
		return err
	}
	if err := s.db.Migrator().CreateTable(&article.Article{}, &transaction.Transaction{}); err != nil {
		return err
	}

	return nil
}

func (s *Server) bootstrapRepositories() {
	s.articleRepo = article.NewRepository(s.db)
	s.transactionRepo = transaction.NewRepository(s.db)

}

func (s *Server) bootstrapServices() {
	s.articleService = article.NewService(s.articleRepo)
	s.transactionService = transaction.NewService(s.transactionRepo)
}

func (s *Server) Run() error {
	logrus.WithField("dsn", s.conf.DSN()).Info("connecting to database")
	var err error

	s.db, err = gorm.Open(
		mysql.Open(s.conf.DSN()), &gorm.Config{},
	)
	if err != nil {
		return err
	}

	if err := s.runMigrations(); err != nil {
		return err
	}

	s.bootstrapRepositories()

	s.bootstrapServices()

	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(
		articlev1connect.ArticleServiceName,
		transactionv1connect.TransactionServiceName,
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle(articlev1connect.NewArticleServiceHandler(s.articleService))
	mux.Handle(transactionv1connect.NewTransactionServiceHandler(s.transactionService))

	if err := http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		return err
	}

	return nil
}
