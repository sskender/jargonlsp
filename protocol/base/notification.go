package base

type NotificationMessage struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
}
