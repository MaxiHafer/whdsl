package emitter

import (
	"context"
	"math/rand"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"

	articlev1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/article/v1/articlev1connect"
	transactionv1 "github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1"
	"github.com/maxihafer/whdsl/pkg/pb/whdsl/transaction/v1/transactionv1connect"
)

func NewEmitterForArticle(
	ctx context.Context,
	conf *ArticleConfig,
	article articlev1connect.ArticleServiceClient,
	transaction transactionv1connect.TransactionServiceClient,
) (*Emitter, error) {

	e := &Emitter{}

	e.ticker = time.NewTicker(conf.Order.Frequency)

	e.articleClient = article
	e.transactionClient = transaction

	logrus.WithFields(
		logrus.Fields{
			"name":            conf.Name,
			"order-frequency": conf.Order.Frequency,
		},
	)

	showDetailsResponse, err := e.articleClient.ShowDetailsForName(
		ctx,
		connect.NewRequest(
			&articlev1.ShowDetailsForNameRequest{Name: conf.Name},
		),
	)

	if err != nil && connect.CodeOf(err) != connect.CodeNotFound{
		return nil, err
	}

	if connect.CodeOf(err) == connect.CodeNotFound {
		createResponse, createErr := e.articleClient.NewArticle(ctx, connect.NewRequest(&articlev1.NewArticleRequest{
			Name:      conf.Name,
			MinAmount: int32(conf.MinimumAmount),
		}))
		if createErr != nil {
			return nil, err
		}

		e.id = createResponse.Msg.GetId()
		logrus.WithField("id", e.id).Info("created new article")
	} else {
		e.id = showDetailsResponse.Msg.GetArticle().GetId()
		logrus.WithField("id", e.id).Info("article already present")
	}

	return e, nil
}

type Emitter struct {
	ticker            *time.Ticker
	articleClient     articlev1connect.ArticleServiceClient
	transactionClient transactionv1connect.TransactionServiceClient
	id                string
	name              string
}

func (e *Emitter) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-e.ticker.C:

			request := connect.NewRequest(
				&transactionv1.NewTransactionRequest{
					ArticleId: e.id,
				},
			)

			count := randInt()

			if count < 0 {
				request.Msg.Type = transactionv1.Transaction_TYPE_OUT
				request.Msg.Count = int32(-count)
			} else {
				request.Msg.Type = transactionv1.Transaction_TYPE_IN
				request.Msg.Count = int32(count)
			}

			logrus.WithFields(
				logrus.Fields{
					"articleId": e.id,
					"count":     request.Msg.GetCount(),
					"type":      request.Msg.GetType().String(),
				},
			).Info("creating transaction")

			_, err := e.transactionClient.NewTransaction(ctx, request)
			if err != nil {
				logrus.Errorf("error while creating transaction: %+v", err)
			}
		}
	}
}

func randInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) - 500
}
