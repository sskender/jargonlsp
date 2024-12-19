package base

type ResponseMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  any    `json:"result"`
}
