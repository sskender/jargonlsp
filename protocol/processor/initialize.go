package processor

import (
	"jargonlsp/protocol/base"
	"log"
)

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeRequestParams struct {
	ProcessId  int         `json:"processId"`
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type InitializeRequest struct {
	base.RequestMessage
	Params *InitializeRequestParams `json:"params"`
}

type ServerCapabilities struct {
	PositionEncoding string `json:"positionEncoding"`
	TextDocumentSync int    `json:"textDocumentSync"`
	HoverProvider    bool   `json:"hoverProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	ServerCapabilities *ServerCapabilities `json:"capabilities"`
	ServerInfo         *ServerInfo         `json:"serverInfo"`
}

func Initialize(requestMessage any) (any, error) {
	message := requestMessage.(*InitializeRequest)

	mid := message.Id
	pid := message.Params.ProcessId

	cname := message.Params.ClientInfo.Name
	cversion := message.Params.ClientInfo.Version

	log.Printf("new client initialize request with id %d and pid %d", mid, pid)
	log.Printf("client is %s v%s", cname, cversion)

	result := InitializeResult{
		ServerCapabilities: &ServerCapabilities{
			PositionEncoding: "utf-16",
			TextDocumentSync: 1,
			HoverProvider:    true,
		},
		ServerInfo: &ServerInfo{
			Name:    base.LSP_SERVER_NAME,
			Version: base.LSP_SERVER_VERSION,
		},
	}

	return &result, nil
}
