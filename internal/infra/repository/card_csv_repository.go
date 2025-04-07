package repository

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"

	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"
)

type CardCSVRepository struct {
	filePath string
}

func NewCardCSVRepository(filePath string) *CardCSVRepository {
	return &CardCSVRepository{filePath}
}

func (r *CardCSVRepository) GetByNumber(number string) (*domain.Card, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
			closeErr := file.Close()
			if err == nil {
					err = closeErr
			}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		
		if strings.HasPrefix(number, row[0]) {
			card := &domain.Card{
				Number:  row[0],
				Name:    row[1],
				Type:    row[2],
				Level:   row[3],
				Country: row[4],
			}
			return card, nil
		}
	}

	return nil, errors.New("card not found")
}