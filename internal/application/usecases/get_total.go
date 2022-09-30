package usecases

import "github.com/rafaelpapastamatiou/imersao-go/internal/application/repositories"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository repositories.OrderRepository
}

func NewGetTotalUseCase(orderRepository repositories.OrderRepository) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{Total: total}, nil
}
