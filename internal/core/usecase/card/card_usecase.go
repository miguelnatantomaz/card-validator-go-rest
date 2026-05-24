package card

import (
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/validator"
)

type CardUsecase struct {
	repo domain.CardRepository
}

func NewCardUsecase(repo domain.CardRepository) *CardUsecase {
	return &CardUsecase{repo}
}

func (u *CardUsecase) GetCardByNumber(number string) (*domain.Card, error) {
	return u.repo.GetByNumber(number)
}

func (u *CardUsecase) ValidateCard(number string) bool {
	return validator.ValidateCard(number)
}
