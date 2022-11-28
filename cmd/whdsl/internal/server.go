package internal

import (
	"log"
	"net/http"
	"os"
	"time"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/maxihafer/whdsl/pkg/article"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

func NewServerFromEnv() (*Server, error){
	s := &Server{}

	conf := NewMariaDBConfigFromEnv()

	s.conf = conf

	return s, nil
}

type Server struct {
	conf *MariadbConfig

	db                 *gorm.DB
	articleRepo        *article.ArticleRepository
	transactionRepo    *article.TransactionRepository
	articleService     *article.ArticleService
	transactionService *article.TransactionService
}

func (s *Server) runMigrations() error {
	if err := s.db.AutoMigrate(&article.Count{}); err != nil {
		return err
	}

	if err := s.db.AutoMigrate(&article.Article{}); err != nil {
		return err
	}

	if err := s.db.AutoMigrate(&article.Transaction{}); err != nil {
		return err
	}

	return nil
}

func (s *Server) bootstrapRepositories() {
	s.articleRepo = article.NewArticleRepository(s.db)
	s.transactionRepo = article.NewTransactionRepository(s.db)

}

func (s *Server) bootstrapServices(){
	s.articleService = article.NewArticleService(s.articleRepo)
	s.transactionService = article.NewTransactionService(s.transactionRepo)
}

func (s *Server) Run() error {
	logrus.WithField("dsn", s.conf.DSN()).Info("connecting to database")
	var err error
	
	s.db, err = gorm.Open(
		mysql.Open(s.conf.DSN()), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					Colorful:      true,
					LogLevel:      logger.Info,
				},
			),
		},
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
		articlev1connect.TransactionServiceName,
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle(articlev1connect.NewArticleServiceHandler(s.articleService))
	mux.Handle(articlev1connect.NewTransactionServiceHandler(s.transactionService))

	if err := http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		return err
	}

	return nil
}
