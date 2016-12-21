package main

import (
	"encoding/json"
	"github.austin.utexas.edu/kriegrj/amqp"
	"github.austin.utexas.edu/kriegrj/hasenachrichtwarteschlange/nutzwert"
	"log"
	"os"
)

func main() {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	nutzwert.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	nutzwert.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"yo",  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	nutzwert.FailOnError(err, "Failed to declare a queue")

	m := nutzwert.Message{"chuckeffinstrong", "Charlie Strong", "former employee", ""}
	body, err := json.Marshal(m)
	nutzwert.FailOnError(err, "Failed to marshal JSON")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	var decodedMsg nutzwert.Message
	err = json.Unmarshal(body, &decodedMsg)
	nutzwert.FailOnError(err, "Failed to unmarshal JSON")
	log.Printf(" [x] Sent %s", decodedMsg)
	nutzwert.FailOnError(err, "Failed to publish a message")
}