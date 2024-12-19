package base

type RequestMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
}
