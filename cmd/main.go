package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelpapastamatiou/imersao-go/internal/application/usecases"
	database "github.com/rafaelpapastamatiou/imersao-go/internal/infra/database/repositories"
	"github.com/rafaelpapastamatiou/imersao-go/internal/presentation/controllers"
	"github.com/rafaelpapastamatiou/imersao-go/pkg/rabbitmq"
)

func main() {
	maxWorkers := 10

	wg := sync.WaitGroup{}

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepository(db)

	calculateFinalPrice := usecases.NewCalculateFinalPriceUseCase(orderRepository)

	getTotalUseCase := usecases.NewGetTotalUseCase(orderRepository)

	getTotalController := controllers.NewGetTotalController(*getTotalUseCase)

	http.HandleFunc("/total", getTotalController.Handle)

	go http.ListenAndServe(":8181", nil)

	conn, ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	defer ch.Close()

	deliveryChannel := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, deliveryChannel)

	wg.Add(maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		go func(j int) {
			defer wg.Done()
			worker(deliveryChannel, calculateFinalPrice, j)
		}(i)
	}

	wg.Wait()
}

func worker(deliveryChannel <-chan amqp.Delivery, useCase *usecases.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryChannel {
		var input usecases.OrderInputDTO

		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			fmt.Println("Error on input unmarshal", err)
		}

		input.Tax = 10.0

		_, err = useCase.Execute(input)

		if err != nil {
			fmt.Println("Error on save order", err)
		}

		msg.Ack(false)

		fmt.Println("Worker", workerId, "processed order", input.ID)
	}
}
