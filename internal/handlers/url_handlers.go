package handler
//internal/hanlers/url_handlers.go
import (
	"net/http"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)
type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var req struct {
		URL   string `json:"url" binding:"required"`
		Alias string `json:"alias,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alias, err := h.service.CreateShortURL(c.Request.Context(), req.URL, req.Alias)
	if err != nil {
		if err == service.ErrEmptyURL {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	host := c.Request.Host
	scheme := "http://"

	c.JSON(http.StatusCreated, gin.H{
		"short_url": scheme + host + "/" + alias,
		"alias":     alias,
	})
}

func (h *URLHandler) RedirectToURL(c *gin.Context) {
	alias := c.Param("alias")
	original, err := h.service.GetOriginalURL(c.Request.Context(), alias)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Redirect(http.StatusFound, original)
}

func (h *URLHandler) DeleteShortURL(c *gin.Context) {
	alias := c.Param("alias")
	if err := h.service.DeleteShortURL(c.Request.Context(), alias); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
