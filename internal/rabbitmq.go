package internal

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitClient is used to keep track of the RabbitMQ connection
type RabbitClient struct {
	// The connection that is used
	conn *amqp.Connection
	// The channel that processes/sends Messages
	ch *amqp.Channel
}

// ConnectRabbitMQ will spawn a Connection
func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	// Set up the Connection to RabbitMQ host using AMQP
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewRabbitMQClient will connect and return a Rabbit client with an open connection
// Accepts an amqp Connection to be reused, to avoid spawning one TCP connection per concurrent client
func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	// Unique, Concurrent Server Channel to process/send messages
	// A good rule of thumb is to always REUSE Conn across applications
	// But spawn a new Channel per routine
	ch, err := conn.Channel()
	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{
		conn: conn,
		ch:   ch,
	}, nil
}

// Close will close the channel
func (rc RabbitClient) Close() error {
	return rc.ch.Close()
}

// CreateQueue will create a new queue based on given cfgs
func (rc RabbitClient) CreateQueue(queueName string, durable, autoDelete bool) error {
	_, err := rc.ch.QueueDeclare(queueName, durable, autoDelete, false, false, nil)
	return err
}

// CreateBinding connects a queue to an exchange using the specified binding rule.
//
// Parameters:
//   - name: The name of the queue.
//   - binding: The binding key for the queue.
//   - exchange: The name of the exchange to which the queue is bound.
//
// The 'nowait' parameter is set to false, which means the channel will wait for the server
// to confirm the binding. If 'nowait' were set to true, the channel would not wait for a
// confirmation and would return immediately.
//
// The 'args' parameter allows for extra arguments, but it's not used in this case.
//
// Returns:
//   - An error if the binding operation fails.
func (rc RabbitClient) CreateBinding(name, binding, exchange string) error {
	return rc.ch.QueueBind(name, binding, exchange, false, nil)
}
