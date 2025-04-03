package adapters

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	albumUsecase *usecase.AlbumUsecase
	cardScraper  *usecase.CardScraper
}

func NewHandler(albumUsecase *usecase.AlbumUsecase, cardScraper *usecase.CardScraper) *Handler {
	return &Handler{albumUsecase, cardScraper}
}

func (h *Handler) GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, h.albumUsecase.GetAllAlbums())
}

func (h *Handler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := h.albumUsecase.GetAlbumByID(id)
	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *Handler) PostAlbum(c *gin.Context) {
	var newAlbum domain.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.albumUsecase.AddAlbum(newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}


func (h *Handler) StartScraping(c *gin.Context) {
	url := os.Getenv("SCRAPER_URL")
	pagesStr := os.Getenv("SCRAPER_PAGES")

	pages, err := strconv.Atoi(pagesStr)
	if err != nil {
		fmt.Println("Erro ao converter SCRAPER_PAGES para int:", err)
		return
	}

	go h.cardScraper.ScrapeCards(url, pages)
	c.JSON(http.StatusOK, gin.H{"message": "Scraping started"})
}