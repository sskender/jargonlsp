package protocol

import (
	"log"
	"os"
)

func ExitNotificationProcessor(_ any) (any, error) {
	defer os.Exit(0)

	log.Printf("Exit request received")

	return nil, nil
}
