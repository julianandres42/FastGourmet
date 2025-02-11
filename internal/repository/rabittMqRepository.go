package repository

import (
	"FastGourmet/internal/core/domain"
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabittMqRepository struct {
	conn *amqp091.Connection
}

func NewRabittMqRepository(conn *amqp091.Connection) *RabittMqRepository {
	return &RabittMqRepository{
		conn: conn,
	}
}

func (r RabittMqRepository) Enqueue(order *domain.Order) error {
	ch, err := r.conn.Channel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
		return err
	}
	err = ch.ExchangeDeclare("orders", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare an exchange", err)
		return err
	}

	jsonBytes, err := json.Marshal(order)
	if err != nil {
		log.Panicf("%s: %s", "Failed to marshal order", err)
		return err
	}

	err = ch.PublishWithContext(ctx, "orders", // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        jsonBytes,
		})

	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
		return err
	}

	log.Printf(" [x] Sent %s", string(jsonBytes))

	return nil

}

func (r RabittMqRepository) Receive(messages chan []byte) {
	//TODO implement me
	ch, err := r.conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	err = ch.ExchangeDeclare("orders", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	q, err := ch.QueueDeclare("", // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	err = ch.QueueBind(q.Name, "", "orders", false, nil)

	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	msgs, err := ch.Consume(q.Name, // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)

	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Got it [x] %s", d.Body)
			messages <- d.Body
		}
	}()

	<-forever
}
