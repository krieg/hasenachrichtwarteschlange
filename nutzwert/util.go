package nutzwert

import (
	"log"
)

// FailOnError logs an error and calls os.Exit(1)
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
