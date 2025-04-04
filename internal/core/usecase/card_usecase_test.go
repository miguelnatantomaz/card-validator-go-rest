package usecase_test

import (
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/domain"
	"github.com/miguelnatantomaz/card-validator-go-rest/internal/core/usecase"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCardRepository struct {
	mock.Mock
}

func (m *MockCardRepository) GetByNumber(number string) (*domain.Card, error) {
	args := m.Called(number)
	if card, ok := args.Get(0).(*domain.Card); ok {
		return card, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestValidateCard_Valid(t *testing.T) {
	repo := new(MockCardRepository)
	uc := usecase.NewCardUsecase(repo)

	validCard := "4539578763621486"

	result := uc.ValidateCard(validCard)

	assert.True(t, result)
}

func TestValidateCard_Invalid(t *testing.T) {
	repo := new(MockCardRepository)
	uc := usecase.NewCardUsecase(repo)

	invalidCard := "1234567890123456"

	result := uc.ValidateCard(invalidCard)

	assert.False(t, result)
}