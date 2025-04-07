package usecase

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"

	"github.com/gocolly/colly"
)

type CardScraper struct{}

func NewCardScraper() *CardScraper {
	return &CardScraper{}
}

func (cs *CardScraper) ScrapeCards(url string, pageSize int) {
	var cards []domain.Card

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       2 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	if err != nil {
		log.Printf("Erro ao limitar colly: %v", err)
		return
	}

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
	})

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			card := domain.Card{
				Number:  el.ChildText("td:nth-child(1)"),
				Name:    el.ChildText("td:nth-child(2)"),
				Type:    el.ChildText("td:nth-child(3)"),
				Level:   el.ChildText("td:nth-child(4)"),
				Country: el.ChildText("td:nth-child(5)"),
			}
			cards = append(cards, card)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("cards.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer func() {
				closeErr := file.Close()
				if err == nil {
						err = closeErr
				}
		}()

		writer := csv.NewWriter(file)
		err = writer.Write([]string{"Number", "Name", "Type", "Level", "Country"})
		if err != nil {
			log.Printf("falha ao inicia writer: %v", err)
				return
		}

		for _, card := range cards {
			err := writer.Write([]string{card.Number, card.Name, card.Type, card.Level, card.Country})
			if err != nil {
				log.Printf("falha ao escrever no CSV: %v", err)
			}
		}

		writer.Flush()
	})

	for i := 1; i <= pageSize; i++ {
		concat_url := fmt.Sprintf("%s%d/", url, i)

		err := c.Visit(concat_url)
		if err != nil {
			log.Println("Error visiting page:", err)
		}
		fmt.Println("Iteração:", concat_url)
	}

	c.Wait()
}
