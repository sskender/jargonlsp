package processor

import "log"

func Shutdown(_ any) (any, error) {

	log.Printf("shutdown request received from client")

	return nil, nil
}
