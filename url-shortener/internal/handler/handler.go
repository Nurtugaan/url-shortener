package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/urlshortener/internal/service"
)

type Handler struct {
	service *service.URLService
}

func NewHandler(service *service.URLService) *Handler {
	return &Handler{service: service}
}

type createURLRequest struct {
	LongURL string `json:"long_url" binding:"required"`
}

func (h *Handler) CreateShortURL(c *gin.Context) {
	var req createURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := h.service.CreateShortURL(req.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, url)
}

func (h *Handler) RedirectToLongURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	longURL, err := h.service.GetLongURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
