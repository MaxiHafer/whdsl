// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3-0.20221101205447-050c4bfe15b5 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for TransactionDirection.
const (
	IN  TransactionDirection = "IN"
	OUT TransactionDirection = "OUT"
)

// Article defines model for Article.
type Article struct {
	CreatedAt time.Time          `bun:",nullzero" json:"createdAt"`
	Id        openapi_types.UUID `bun:",pk,type:varchar(36)" json:"id"`
	MinAmount int                `json:"minAmount"`
	Name      string             `json:"name"`
	UpdatedAt time.Time          `bun:",nullzero" json:"updatedAt"`
}

// ArticleProperties defines model for ArticleProperties.
type ArticleProperties struct {
	CreatedAt *time.Time          `bun:",nullzero" json:"createdAt,omitempty"`
	Id        *openapi_types.UUID `bun:",pk,type:varchar(36)" json:"id,omitempty"`
	MinAmount *int                `json:"minAmount,omitempty"`
	Name      *string             `json:"name,omitempty"`
	UpdatedAt *time.Time          `bun:",nullzero" json:"updatedAt,omitempty"`
}

// Transaction defines model for Transaction.
type Transaction struct {
	Amount    *int                  `json:"amount,omitempty"`
	ArticleId *openapi_types.UUID   `json:"articleId,omitempty"`
	CreatedAt *time.Time            `json:"createdAt,omitempty"`
	Direction *TransactionDirection `json:"direction,omitempty"`
	Id        *openapi_types.UUID   `json:"id,omitempty"`
	UpdatedAt *time.Time            `json:"updatedAt,omitempty"`
}

// TransactionDirection defines model for Transaction.Direction.
type TransactionDirection string

// PostArticlesJSONRequestBody defines body for PostArticles for application/json ContentType.
type PostArticlesJSONRequestBody = Article

// PostTransactionsJSONRequestBody defines body for PostTransactions for application/json ContentType.
type PostTransactionsJSONRequestBody = Transaction
