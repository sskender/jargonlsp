package protocol

import (
	"encoding/json"
	"fmt"
)

func encodeServerResponse(response any) ([]byte, error) {
	content, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

	return []byte(result), nil
}

func decodeClientRequest(content []byte, request any) error {
	return json.Unmarshal(content, request)
}

func getRequestProcessor(method string) (func(any) (any, error), any, error) {
	switch method {

	case METHOD_INITIALIZE:
		return InitializeRequestProcessor, &InitializeRequest{}, nil

	case METHOD_INITIALIZED:
		return InitializedNotificationProcessor, &NotificationMessage{}, nil

	case METHOD_SHUTDOWN:
		return ShutdownRequestProcessor, &RequestMessage{}, nil

	case METHOD_EXIT:
		return ExitNotificationProcessor, &NotificationMessage{}, nil

	default:
		return nil, nil, fmt.Errorf("unknown request message method: '%s'", method)

	}
}

func HandleClientRequest(content []byte) ([]byte, error) {
	requestMessage := RequestMessage{}

	err := decodeClientRequest(content, &requestMessage)
	if err != nil {
		return nil, err
	}

	fRequestProcessor, requestMessageType, err := getRequestProcessor(requestMessage.Method)
	if err != nil {
		return nil, err
	}

	err = decodeClientRequest(content, requestMessageType)
	if err != nil {
		return nil, err
	}

	result, err := fRequestProcessor(requestMessageType)
	if err != nil {
		return nil, err
	}

	responseMessage := ResponseMessage{
		JsonRPC: requestMessage.JsonRPC,
		Id:      requestMessage.Id,
		Result:  result,
	}

	response, err := encodeServerResponse(responseMessage)
	if err != nil {
		return nil, err
	}

	return response, err
}
