package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
			log.Println("Error closing amqp connection")
		}
	}(connectRabbitMQ)

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
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

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"proapi-test", // queue name
		true,          // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		panic(err)
	}

	// Create a new Fiber instance.
	app := fiber.New()

	// Add middleware.
	app.Use(
		logger.New(), // add simple logger
	)

	// Add route.
	app.Get("/send", func(c *fiber.Ctx) error {
		// Create a message to publish.
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}

		// Attempt to publish a message to the queue.
		if err := channelRabbitMQ.Publish(
			"",            // exchange
			"proapi-test", // queue name
			false,         // mandatory
			false,         // immediate
			message,       // message to publish
		); err != nil {
			return err
		}

		return nil
	})

	// Start Fiber API server.
	log.Fatal(app.Listen(":3000"))
}
