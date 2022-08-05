package main

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer func(connectRabbitMQ *amqp.Connection) {
		err := connectRabbitMQ.Close()
		if err != nil {
			log.Println("Error closing amqp connection!")
		}
	}(connectRabbitMQ)

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer func(channelRabbitMQ *amqp.Channel) {
		err := channelRabbitMQ.Close()
		if err != nil {
			log.Println("Error closing channel connection!")
		}
	}(channelRabbitMQ)

	// Subscribing to proapi-test for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"proapi-test", // queue name
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-forever
}
