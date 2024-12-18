package processor

import "log"

func Initialized(_ any) (any, error) {

	log.Println("new client initialized successfully")

	return nil, nil
}
