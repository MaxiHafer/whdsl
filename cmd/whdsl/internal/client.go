package internal

import (
	"net/http"

	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
)

func NewClientFromEnv() *Client {
	c := &Client{}

	c.conf = NewClientConfigFromEnv()

	return c
}

type Client struct {
	conf *ClientConfig

	article articlev1connect.ArticleServiceClient
	transaction articlev1connect.TransactionServiceClient
}

func (c *Client) Run() error {
	httpClient := http.DefaultClient

	c.article = articlev1connect.NewArticleServiceClient(httpClient, c.conf.DSN())
	c.transaction = articlev1connect.NewTransactionServiceClient(httpClient, c.conf.DSN())

	return nil
}
