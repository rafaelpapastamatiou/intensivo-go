package entities_test

import (
	"testing"

	"github.com/rafaelpapastamatiou/imersao-go/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGivenAndEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entities.Order{}

	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAndEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entities.Order{
		ID: "123",
	}

	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAndEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entities.Order{
		ID:    "123",
		Price: 10,
	}

	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	_, err := entities.NewOrder("123", 10, 2)

	assert.NoError(t, err)
}

func TestGivenValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetFinalPriceProperty(t *testing.T) {
	order, err := entities.NewOrder("123", 10, 2)

	assert.NoError(t, err)

	err = order.CalculateFinalPrice()

	assert.NoError(t, err)

	assert.Equal(t, 12.0, order.FinalPrice)
}
