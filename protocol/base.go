package protocol

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

// TODO implement Error

// TODO implement Notification
