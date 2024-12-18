package base

const (
	LSP_SERVER_NAME    = "JargonLSP"
	LSP_SERVER_VERSION = "0.1.0"
)

type RequestMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
}

type ResponseMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  any    `json:"result"`
}

type NotificationMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
}

// TODO implement Error
