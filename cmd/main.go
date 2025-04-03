package main

import (
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/adapters"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/usecase"
)

func main() {
	albumUsecase := usecase.NewAlbumUsecase()
	cardScraper := usecase.NewCardScraper()

	handler := adapters.NewHandler(albumUsecase, cardScraper)
	router := adapters.SetupRouter(handler)

	router.Run(":8080")
}