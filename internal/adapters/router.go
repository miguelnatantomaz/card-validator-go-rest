package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/albums", handler.GetAlbums)
	r.GET("/albums/:id", handler.GetAlbumByID)
	r.POST("/albums", handler.PostAlbum)
	r.GET("/scrape", handler.StartScraping)
	r.GET("/cards/:number", handler.GetCardByNumber)

	return r
}