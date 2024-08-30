package main

import (
	"github.com/Pomog/rabbitmq-sqs-sns-go/internal"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("admin", "password", "localhost:5672", "customers")

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

	if err := client.CreateQueue("customers_created", true, false); err != nil {
		panic(err)
	}
	if err := client.CreateQueue("customers_test", false, true); err != nil {
		panic(err)
	}

	// Create binding between the customer_events exchange and the customers-created queue
	if err := client.CreateBinding("customers_created", "customers.created.*", "customer_events"); err != nil {
		panic(err)
	}
	// Create binding between the customer_events exchange and the customers-test queue
	if err := client.CreateBinding("customers_test", "customers.*", "customer_events"); err != nil {
		panic(err)
	}

	time.Sleep(30 * time.Second)

	log.Println(client)
}
