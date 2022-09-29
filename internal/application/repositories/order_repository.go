package repositories

import "github.com/rafaelpapastamatiou/imersao-go/internal/domain/entities"

type OrderRepository interface {
	Save(o *entities.Order) error
	GetTotal() (int, error)
}
