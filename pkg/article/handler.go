package article

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

type Body struct {
	Name          string `json:"name" example:"Kühlschrank"`
	MinimumAmount int    `json:"min_amount" example:"100"`
}

func (h *Handler) RegisterRoutes() {
	h.r.GET("/", h.OnListArticles)
	h.r.GET("/:id", h.OnGetArticle)
	h.r.POST("/", h.OnCreateArticle)
	h.r.PUT("/:id", h.OnUpdateArticle)
	h.r.DELETE("/:id", h.OnDeleteArticle)
}

func (h *Handler) OnGetArticle(c *gin.Context) {
	article, err := h.repo.GetArticleByID(c, c.Param("id"))
	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no article found for id: %s", c.Param("id")))
	}
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed loading article for id: %s", c.Param("id")))
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) OnListArticles(c *gin.Context) {
	articles, err := h.repo.ListArticles(c)
	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no articles present in repo"))
	}

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed loading articles from repo"))
	}

	c.JSON(http.StatusOK, articles)
}

func (h *Handler) OnCreateArticle(c *gin.Context) {
	body := Body{}

	if err := c.Bind(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.Wrapf(err,"failed binding request body"))
	}

	article := &Model{
		ID:            uuid.NewString(),
		Name:          body.Name,
		MinimumAmount: body.MinimumAmount,
	}

	if err := h.repo.Store(c, article); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Errorf("failed storing article in respository"))
	}

	c.JSON(http.StatusCreated, article.ID)
}

func (h *Handler) OnUpdateArticle(c *gin.Context) {
	body := Body{}
	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.Wrapf(err, "failed binding request body"))
	}

	article, err := h.repo.GetArticleByID(c, c.Param("id"))
	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no article found for id: %s",c.Param("id")))
	}

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed loading article from repo"))
	}

	article.Name = body.Name
	article.MinimumAmount = body.MinimumAmount

	if err := h.repo.Store(c, article); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed while storing article in repository"))
		return
	}

	c.JSON(http.StatusOK, article.ID)
}


func (h *Handler) OnDeleteArticle(c *gin.Context) {
	err := h.repo.DeleteByID(c, c.Param("id"))

	if errors.Is(err, sql.ErrNoRows) {
		_ = c.AbortWithError(http.StatusNotFound, errors.Errorf("no article found for id: %s",c.Param("id")))
	}

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Wrapf(err, "failed while deleting article from repo"))
	}

	c.Status(http.StatusOK)
}
