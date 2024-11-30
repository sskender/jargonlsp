package protocol

import "log"

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeRequestParams struct {
	ProcessId  int         `json:"processId"`
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type InitializeRequest struct {
	RequestMessage
	Params *InitializeRequestParams `json:"params"`
}

type ServerCapabilities struct {
	PositionEncoding string `json:"positionEncoding"`
	TextDocumentSync int    `json:"textDocumentSync"`
	HoverProvider    bool   `json:"hoverProvider"`
	// TODO idea: colors for known words?
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	ServerCapabilities *ServerCapabilities `json:"capabilities"`
	ServerInfo         *ServerInfo         `json:"serverInfo"`
}

func InitializeRequestProcessor(requestMessage any) (any, error) {
	message := requestMessage.(*InitializeRequest)

	log.Printf("new client initialize request with id %d and pid %d", message.Id, message.Params.ProcessId)
	log.Printf("client is %s v%s", message.Params.ClientInfo.Name, message.Params.ClientInfo.Version)

	// TODO implement logic to save client state

	result := InitializeResult{
		ServerCapabilities: &ServerCapabilities{
			PositionEncoding: "utf-16",
			TextDocumentSync: 1,
			HoverProvider:    true,
		},
		ServerInfo: &ServerInfo{
			Name:    "JargonLSP", // TODO dont hardcode
			Version: "0.1.0",     // TODO dont hardcode
		},
	}

	return &result, nil
}
