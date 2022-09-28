package transaction

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"whdsl/pkg/api"
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
}

func (h *Handler) OnGetTransactionByID(c *gin.Context) {
	transaction, err := h.repo.GetTransactionByID(c, c.Param("id"))
	if errors.Is(err, sql.ErrNoRows) {
		api.ErrNotFound(c)
	}
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) OnListTransactions(c *gin.Context) {
	transactions, err := h.repo.ListTransactions(c)
	if errors.Is(err, sql.ErrNoRows) {
		api.ErrNotFound(c)
	}

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}
