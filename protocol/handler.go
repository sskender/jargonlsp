package protocol

import (
	"encoding/json"
	"log"
)

const (
	METHOD_INITIALIZE_REQUEST = "initialize"
)

type RequestMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	// TODO Params as string to parse later
}

func decodeRequestMessage(content []byte) (*RequestMessage, error) {
	var requestMessage RequestMessage = RequestMessage{}

	err := json.Unmarshal(content, &requestMessage)
	if err != nil {
		return nil, err
	}

	return &requestMessage, nil
}

// TODO implement encode message

func handleUnknownMethod(requestMessage *RequestMessage) {
	log.Printf("unknown request method %s", requestMessage.Method)
}

func handleInitializeRequestMethod(requestMessage *RequestMessage) {
	log.Println("handling initizalie request")
	log.Println(requestMessage)
}

func HandleRequestMessage(content []byte) error {
	requestMessage, err := decodeRequestMessage(content)
	if err != nil {
		return err
	}

	log.Println("THE PARSED MESSAGE:")
	log.Println(requestMessage)
	log.Println("END PARSED MESSAGE:")

	// TODO where is the response handled - need to pass it to server so it can forward it

	switch requestMessage.Method {

	case METHOD_INITIALIZE_REQUEST:
		handleInitializeRequestMethod(requestMessage)

	default:
		handleUnknownMethod(requestMessage)

	}

	return nil
}
