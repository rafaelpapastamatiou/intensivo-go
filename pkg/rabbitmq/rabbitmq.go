package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, nil, err
	}

	ch.Qos(100, 0, false)

	return conn, ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
