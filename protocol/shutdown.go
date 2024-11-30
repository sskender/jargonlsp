package protocol

import "log"

func ShutdownRequestProcessor(requestMessage any) (any, error) {
	message := requestMessage.(*RequestMessage)

	log.Printf("Shutdown request received from client with id %d", message.Id)

	return nil, nil
}
