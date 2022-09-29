package main

import (
	"encoding/json"
	"math/rand"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelpapastamatiou/imersao-go/pkg/rabbitmq"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrder() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
	}
}

func Notify(ch *amqp.Channel, order Order) error {
	body, err := json.Marshal(order)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	conn, ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	defer ch.Close()

	for i := 0; i < 100; i++ {
		order := GenerateOrder()

		err := Notify(ch, order)

		if err != nil {
			panic(err)
		}
	}
}
