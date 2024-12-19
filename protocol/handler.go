package protocol

import (
	"encoding/json"
	"fmt"
	"jargonlsp/protocol/base"
	"jargonlsp/protocol/processor"
	"log"
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

func getRequestProcessor(method string) (func(any) (any, error), any) {
	switch method {

	case METHOD_INITIALIZE:
		return processor.Initialize, &processor.InitializeRequest{}

	case METHOD_INITIALIZED:
		return processor.Initialized, &base.NotificationMessage{}

	case METHOD_SHUTDOWN:
		return processor.Shutdown, &base.RequestMessage{}

	case METHOD_EXIT:
		return processor.Exit, &base.NotificationMessage{}

	case METHOD_TEXT_DOC_OPEN:
		return processor.DocumentDidOpen, &processor.DidOpenTextDocumentNotification{}

	case METHOD_TEXT_DOC_CHANGE:
		return processor.DocumentDidChange, &processor.DidChangeTextDocumentNotification{}

	case METHOD_TEXT_DOC_SAVE:
		return processor.DocumentDidSave, &processor.DidSaveTextDocumentNotification{}

	case METHOD_TEXT_DOC_CLOSE:
		return processor.DocumentDidClose, &processor.DidCloseTextDocumentNotification{}

	default:
		return nil, nil

	}
}

func HandleClientRequest(content []byte) ([]byte, error) {
	requestMessage := base.RequestMessage{}

	err := decodeClientRequest(content, &requestMessage)
	if err != nil {
		return nil, err
	}

	fRequestProcessor, tRequestMessage := getRequestProcessor(requestMessage.Method)

	if fRequestProcessor == nil || tRequestMessage == nil {
		log.Printf("unknown request message method: '%s'", requestMessage.Method)
		return nil, nil
	}

	err = decodeClientRequest(content, tRequestMessage)
	if err != nil {
		return nil, err
	}

	fResponse, err := fRequestProcessor(tRequestMessage)
	if err != nil {
		return nil, err
	}

	if fResponse == nil {
		return nil, nil
	}

	responseMessage := base.ResponseMessage{
		JsonRPC: requestMessage.JsonRPC,
		Id:      requestMessage.Id,
		Result:  fResponse,
	}

	response, err := encodeServerResponse(responseMessage)
	if err != nil {
		return nil, err
	}

	return response, err
}
