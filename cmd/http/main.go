package main

import (
	"database/sql"
	"net/http"

	"github.com/rafaelpapastamatiou/imersao-go/internal/application/usecases"
	database "github.com/rafaelpapastamatiou/imersao-go/internal/infra/database/repositories"
	"github.com/rafaelpapastamatiou/imersao-go/internal/presentation/controllers"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepository(db)

	getTotalUseCase := usecases.NewGetTotalUseCase(orderRepository)

	getTotalController := controllers.NewGetTotalController(*getTotalUseCase)

	http.HandleFunc("/total", getTotalController.Handle)

	http.ListenAndServe(":8181", nil)
}
