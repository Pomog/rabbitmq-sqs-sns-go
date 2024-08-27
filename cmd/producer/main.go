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

	time.Sleep(30 * time.Second)

	log.Println(client)
}
