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
	cardUseCase  *usecase.CardUsecase
}

func NewHandler(
	albumUsecase *usecase.AlbumUsecase, 
	cardScraper *usecase.CardScraper,
	cardUsecase *usecase.CardUsecase,
	) *Handler {
	return &Handler{
		albumUsecase, 
		cardScraper,
		cardUsecase,
	}
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

func (h *Handler) GetCardByNumber(c *gin.Context) {
	id := c.Param("number")
	card, err := h.cardUseCase.GetCardByNumber(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Card not found"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func (h *Handler) ValidateCard(c *gin.Context) {
	number := c.Param("number")
	isValid := h.cardUseCase.ValidateCard(number)

	c.JSON(http.StatusOK, gin.H{
		"number": number,
		"valid": isValid,
	})
}