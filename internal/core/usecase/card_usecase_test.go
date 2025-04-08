package usecase_test

import (
	"errors"

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

func TestGetCardByNumber_Success(t *testing.T) {
	repo := new(MockCardRepository)
	uc := usecase.NewCardUsecase(repo)
	
	expectedCard := &domain.Card{
		Number:  "4539578763621486",
		Name:    "Visa",
		Type:    "Credit",
		Level:   "Gold",
		Country: "US",
	}
	
	repo.On("GetByNumber", "4539578763621486").Return(expectedCard, nil)
	
	card, err := uc.GetCardByNumber("4539578763621486")
	
	assert.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, expectedCard, card)
	repo.AssertExpectations(t)
}

func TestGetCardByNumber_NotFound(t *testing.T) {
	repo := new(MockCardRepository)
	uc := usecase.NewCardUsecase(repo)
	
	expectedErr := errors.New("card not found")
	repo.On("GetByNumber", "1234567890123456").Return(nil, expectedErr)
	
	card, err := uc.GetCardByNumber("1234567890123456")
	
	assert.Error(t, err)
	assert.Nil(t, card)
	assert.Equal(t, expectedErr, err)
	repo.AssertExpectations(t)
}

func TestGetCardByNumber_RepositoryError(t *testing.T) {
	repo := new(MockCardRepository)
	uc := usecase.NewCardUsecase(repo)
	
	expectedErr := errors.New("csv not foud")
	repo.On("GetByNumber", "4539578763621486").Return(nil, expectedErr)
	
	card, err := uc.GetCardByNumber("4539578763621486")
	
	assert.Error(t, err)
	assert.Nil(t, card)
	assert.Equal(t, expectedErr, err)
	repo.AssertExpectations(t)
}