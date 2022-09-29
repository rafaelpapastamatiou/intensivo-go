package usecases

import (
	"github.com/rafaelpapastamatiou/imersao-go/internal/application/repositories"
	"github.com/rafaelpapastamatiou/imersao-go/internal/domain/entities"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository repositories.OrderRepository
}

func NewCalculateFinalPriceUseCase(orderRepository repositories.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entities.NewOrder(input.ID, input.Price, input.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
