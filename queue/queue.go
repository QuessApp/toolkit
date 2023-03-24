package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// Connect connects to MQ.
func Connect(connUri string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(connUri)

	if err != nil {
		log.Fatalln(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully connected to MQ instance at %s \n", connUri)

	return conn, ch
}

// Consume consumes a message from a specif queue.
func Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	return msgs, err
}

// Publish publishes a message to a specif queue.
func Publish(ch *amqp.Channel, exchange, queueName, msg string) error {
	err := ch.Publish(
		exchange,
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	return err
}
