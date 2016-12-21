package nutzwert

import (
	"log"
)

// Message represents employee data
type Message struct {
	EID         string
	Name        string
	Affiliation string
	Org         string
}

// FailOnError logs an error and calls os.Exit(1)
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
