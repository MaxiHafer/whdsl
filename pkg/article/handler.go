package article

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	h.r.POST("/", h.OnCreateArticle)
}

// OnGetArticleByID godoc
// @Summary     Show an article
// @Description gets an article by ID
// @Tags        articles
// @Accept      json
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
// @Accept      json
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

type addArticle struct {
	Name          string `json:"name" example:"kebab"`
	MinimumAmount int    `json:"min_amount" example:"1"`
}

// OnCreateArticle godoc
// @Summary     Create article
// @Description creates article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     addArticle true "Add Article"
// @Success     201     {object} string
// @Failure     400     {string} string
// @Router      /articles [post]
func (h *Handler) OnCreateArticle(c *gin.Context) {
	addArticle := &addArticle{}
	if err := c.BindJSON(addArticle); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	article := &Article{
		ID:            uuid.NewString(),
		Name:          addArticle.Name,
		MinimumAmount: addArticle.MinimumAmount,
	}

	if err := h.repo.Store(c, article); err != nil {
		api.ErrInternal(c)
		return
	}

	c.JSON(http.StatusCreated, article.ID)
}

// OnUpdateArticle godoc
// @Summary     Update article
// @Description updates article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body updateArticle true "Update Article"
// @Success     200     {s}
