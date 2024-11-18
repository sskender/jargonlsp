package protocol

import (
	"encoding/json"
	"fmt"
	"log"
)

func EncodeResponseMessage(responseMessage any) ([]byte, error) {
	content, err := json.Marshal(responseMessage)
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

	log.Println("TOTAL BYTES TO WRITE CHECK:", len(content))

	return []byte(result), nil
}

func DecodeRequestMessage(content []byte, requestMessage any) error {
	return json.Unmarshal(content, requestMessage)
}

func GetRequestProcessor(method string) (func(any) (any, error), any, error) {
	switch method {

	case METHOD_INITIALIZE:
		return InitializeRequestProcessor, &InitializeRequest{}, nil

	case METHOD_INITIALIZED:
		return InitializedNotificationProcessor, &RequestMessage{}, nil

	default:
		return nil, nil, fmt.Errorf("unknown request message method: '%s", method)

	}
}

func HandleRequestMessage(content []byte) ([]byte, error) {

	requestMessage := RequestMessage{}
	err := DecodeRequestMessage(content, &requestMessage)
	if err != nil {
		return nil, err
	}

	log.Println("THE PARSED MESSAGE:")
	log.Println(requestMessage)
	log.Println("END PARSED MESSAGE:")

	fRequestProcessor, requestMessageType, err := GetRequestProcessor(requestMessage.Method)
	if err != nil {
		return nil, err
	}

	err = DecodeRequestMessage(content, requestMessageType)
	if err != nil {
		return nil, err
	}

	log.Println("AFTER DECODE")
	log.Println(requestMessageType)

	result, err := fRequestProcessor(requestMessageType)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	responseMessage := ResponseMessage{
		JsonRPC: requestMessage.JsonRPC,
		Id:      requestMessage.Id,
		Result:  result,
	}

	log.Println("response")
	log.Println(responseMessage)

	response, err := EncodeResponseMessage(responseMessage)
	if err != nil {
		return nil, err
	}

	return response, err
}
