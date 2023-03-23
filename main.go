package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	defer conn.Close()

	log.Info().Msg("Connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().Msg("Successfully Published Message to Queue")
}
