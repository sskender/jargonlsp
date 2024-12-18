package processor

import (
	"log"
	"os"
)

func Exit(_ any) (any, error) {
	defer os.Exit(0)

	log.Printf("exit request received")

	// TODO error if not initialized before

	return nil, nil
}
