package main

import (
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/adapters"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/usecase"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/infra/repository"
)

func main() {
	albumUsecase := usecase.NewAlbumUsecase()
	cardScraper := usecase.NewCardScraper()

	cardRepo := repository.NewCardCSVRepository("cards.csv")
	cardUsecase := usecase.NewCardUsecase(cardRepo)

	handler := adapters.NewHandler(albumUsecase, cardScraper, cardUsecase)
	router := adapters.SetupRouter(handler)

	router.Run(":8080")
}