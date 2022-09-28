package article

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
	h.r.GET("/", h.OnListArticles)
	h.r.GET("/:id", h.OnGetArticleByID)
}

// OnGetArticleByID godoc
// @Summary     Show an article
// @Description gets an article by ID
// @Tags        articles
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} Article
// @Failure     404 {string} string
// @Router      /articles/{id} [get]
func (h *Handler) OnGetArticleByID(c *gin.Context) {
	article, err := h.repo.GetArticleByID(c, c.Param("id"))
	if errors.Is(err, sql.ErrNoRows) {
		api.ErrNotFound(c)
	}
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, article)
}

// OnListArticles godoc
// @Summary     List articles
// @Description gets accounts
// @Tags        articles
// @Produce     json
// @Success     200 {array}  Article
// @Failure     404 {string} string
// @Router      /articles [get]
func (h *Handler) OnListArticles(c *gin.Context) {
	articles, err := h.repo.ListArticles(c)
	if errors.Is(err, sql.ErrNoRows) {
		api.ErrNotFound(c)
	}

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, articles)
}
