package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/adapters"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/usecase/album"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/usecase/card"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/infra/repository"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/infra/scraper"
)

func main() {
	// Carrega o arquivo .env se ele existir na raiz do projeto
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado ou não pôde ser carregado.")
	}

	albumUsecase := album.NewAlbumUsecase()
	cardScraper := scraper.NewCardScraper()

	cardRepo := repository.NewCardCSVRepository("cards.csv")
	cardUsecase := card.NewCardUsecase(cardRepo)

	handler := adapters.NewHandler(albumUsecase, cardScraper, cardUsecase)
	router := adapters.SetupRouter(handler)

	err := router.Run(":8080")

	if err != nil {
		log.Printf("Erro ao iniciar servidor: %v", err)
		return
	}
}