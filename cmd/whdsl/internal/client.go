package internal

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/maxihafer/whdsl/pkg/emitter"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1/transactionv1connect"
)

func NewClientFromEnv(emitterConfig *emitter.Config) (*Client, error) {
	c := &Client{}
	
	conf := NewClientConfigFromEnv()

	httpClient := http.DefaultClient
	c.article = articlev1connect.NewArticleServiceClient(httpClient, conf.DSN())
	c.transaction = transactionv1connect.NewTransactionServiceClient(httpClient, conf.DSN())
	c.emitterMap = make(map[string]*emitter.Emitter)
	
	c.emitterConfig = emitterConfig

	return c, nil
}

type Client struct {
	emitterConfig *emitter.Config

	article     articlev1connect.ArticleServiceClient
	transaction transactionv1connect.TransactionServiceClient

	emitterMap map[string]*emitter.Emitter
}

func (c *Client) Run(ctx context.Context) error {

	errGroup := errgroup.Group{}
	
	for _, articleConfig := range c.emitterConfig.Articles {
	
		articleEmitter, err := emitter.NewEmitterForArticle(ctx, articleConfig, c.article, c.transaction)
		if err != nil {
			return errors.Wrapf(err,"error while creating emitter for article '%s",articleConfig.Name)
		}
		
		errGroup.Go(
			func() error {
				return articleEmitter.Start(ctx)
			},
		)
	}

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
}
