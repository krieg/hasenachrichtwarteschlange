package nutzwert

import (
	"github.austin.utexas.edu/kriegrj/amqp"
	"os"
)

// Herausholen retrieves messages from the queue and returns them to the caller.
func Herausholen() <-chan amqp.Delivery {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"affChange", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	FailOnError(err, "Failed to register a consumer")

	return msgs
}
