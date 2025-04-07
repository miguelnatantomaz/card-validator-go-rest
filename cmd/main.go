package main

import (
	"log"

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

	err := router.Run(":8080")

	if err != nil {
		log.Printf("Erro ao limitar colly: %v", err)
		return
	}
}