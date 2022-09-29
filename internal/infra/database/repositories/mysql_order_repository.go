package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelpapastamatiou/imersao-go/internal/domain/entities"
)

type MySqlOrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *MySqlOrderRepository {
	return &MySqlOrderRepository{Db: db}
}

func (r *MySqlOrderRepository) Save(order *entities.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *MySqlOrderRepository) GetTotal() (int, error) {
	var total int

	err := r.Db.QueryRow("SELECT count(*) from orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
