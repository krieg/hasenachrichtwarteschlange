package main

import (
	"github.austin.utexas.edu/kriegrj/hasenachrichtwarteschlange/nutzwert"
	"log"
)

func main() {
	msgs := nutzwert.Herausholen()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
