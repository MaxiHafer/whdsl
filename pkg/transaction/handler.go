package transaction

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewHandler(g *gin.RouterGroup, repo *Repository) *Handler {
	return &Handler{
		repo: repo,
		r:    g,
	}
}

type Handler struct {
	repo *Repository
	r    *gin.RouterGroup
}

func (h *Handler) RegisterRoutes() {
	h.r.GET("/", h.OnListTransactions)
	h.r.GET("/:id", h.OnGetTransactionByID)
	h.r.POST("/", h.OnCreateTransaction)
	h.r.PUT("/:id", h.OnUpdateTransaction)
	h.r.DELETE("/:id", h.OnDeleteTransaction)
}

func (h *Handler) OnGetTransactionByID(c *gin.Context) {
	transaction, err := h.repo.GetTransactionByID(c, c.Param("id"))
	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no transaction found for id: %s",c.Param("id")))
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed while loading transaction from repo"))
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) OnListTransactions(c *gin.Context) {
	transactions, err := h.repo.ListTransactions(c)
	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no transactions in repo"))
	}

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed while loading transactions from repo"))
		return
	}

	c.JSON(http.StatusOK, transactions)
}

type Body struct {
	ArticleID string `json:"article_id" example:"c2144c9d-1b61-49b9-a028-b01029106ca7"`
	Direction Direction `json:"direction" example:"0"`
	Amount    int `json:"amount" example:"100"`
}

func (h *Handler) OnCreateTransaction(c *gin.Context) {
	body := Body{}

	if err := c.Bind(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.Wrapf(err,"failed binding request body"))
	}

	transaction := &Model{
		ID:        uuid.NewString(),
		ArticleID: body.ArticleID,
		Direction: body.Direction,
		Amount:    body.Amount,
	}

	if err := h.repo.Store(c, transaction); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Errorf("failed storing transaction in respository"))
	}

	c.JSON(http.StatusCreated, transaction.ID)
}

func (h *Handler) OnUpdateTransaction(c *gin.Context) {
	body := Body{}
	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.Wrapf(err, "failed binding request body"))
	}

	transaction, err := h.repo.GetTransactionByID(c, c.Param("id"))
	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no transaction found for id: %s",c.Param("id")))
	}

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed loading transaction from repo"))
	}

	transaction.Direction = body.Direction
	transaction.ArticleID = body.ArticleID
	transaction.Amount = body.Amount

	if err := h.repo.Store(c, transaction); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed while storing transaction in repository"))
		return
	}

	c.JSON(http.StatusOK, transaction.ID)
}

func (h *Handler) OnDeleteTransaction(c *gin.Context) {
	err := h.repo.DeleteByID(c, c.Param("id"))

	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no transaction found for id: %s",c.Param("id")))
	}

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed while deleting transaction from repo"))
	}

	c.Status(http.StatusOK)
}
