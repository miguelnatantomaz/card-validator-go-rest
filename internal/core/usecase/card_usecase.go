package usecase

import (
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/validator"
)

// type CardUsecase struct{
// 	repo repository.CardCSVRepository
// }

// func NewCardUsecase(repo repository.CardCSVRepository) *CardUsecase {
// 	return &CardUsecase{repo}
// }

type CardRepository interface {
	GetByNumber(number string) (*domain.Card, error)
}

type CardUsecase struct {
	repo CardRepository
}

func NewCardUsecase(repo CardRepository) *CardUsecase {
	return &CardUsecase{repo}
}

func (cs *CardUsecase) GetCardByNumber(number string) (*domain.Card, error) {
	return cs.repo.GetByNumber(number)
}

func (u *CardUsecase) ValidateCard(number string) bool {
	return validator.ValidateCard(number)
}