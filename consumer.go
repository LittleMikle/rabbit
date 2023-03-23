package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Info().Msgf("Recieved Message: %s\n", d.Body)
		}
	}()

	log.Info().Msg("[*] - Waiting for messages")
	<-forever
}
