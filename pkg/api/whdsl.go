package api

import (
	"context"
)

//go:generate oapi-codegen --config=types.cfg.yaml ../../openapi.yaml
//go:generate oapi-codegen --config=server.cfg.yaml ../../openapi.yaml

var _ StrictServerInterface = &Service{}

func NewService() *Service {
	return &Service{}
}

type Service struct {
}

func (s Service) GetArticles(ctx context.Context, request GetArticlesRequestObject) (GetArticlesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) PostArticles(ctx context.Context, request PostArticlesRequestObject) (PostArticlesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteArticlesId(ctx context.Context, request DeleteArticlesIdRequestObject) (DeleteArticlesIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetArticlesId(ctx context.Context, request GetArticlesIdRequestObject) (GetArticlesIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) PutArticlesId(ctx context.Context, request PutArticlesIdRequestObject) (PutArticlesIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetTransactions(ctx context.Context, request GetTransactionsRequestObject) (GetTransactionsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) PostTransactions(ctx context.Context, request PostTransactionsRequestObject) (PostTransactionsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteTransactionsId(ctx context.Context, request DeleteTransactionsIdRequestObject) (DeleteTransactionsIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetTransactionsId(ctx context.Context, request GetTransactionsIdRequestObject) (GetTransactionsIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) PutTransactionsId(ctx context.Context, request PutTransactionsIdRequestObject) (PutTransactionsIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

