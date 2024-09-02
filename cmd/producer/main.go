package main

import (
	"context"
	"github.com/Pomog/rabbitmq-sqs-sns-go/internal"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("percy", "secret", "localhost:5672", "customers")

	if err != nil {
		panic(err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Critical Error")
		}
	}(conn)
	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer func(client internal.RabbitClient) {
		err := client.Close()
		if err != nil {
			log.Fatal("Critical Error")
		}
	}(client)

	// Create context to manage timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Create customer from sweden
	for i := 0; i < 10; i++ {
		if err := client.Send(ctx, "customer_events", "customers.created.se", amqp.Publishing{
			ContentType:  "text/plain",    // The payload we send is plaintext, could be JSON or others.
			DeliveryMode: amqp.Persistent, // This tells rabbitMQ that this message should be Saved if no resources accepts it before a restart (durable)
			Body:         []byte("An cool message between services"),
		}); err != nil {
			panic(err)
		}
	}

	log.Println(client)
}
