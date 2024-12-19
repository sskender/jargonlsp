package protocol

const (
	METHOD_INITIALIZE  = "initialize"
	METHOD_INITIALIZED = "initialized"
	METHOD_SHUTDOWN    = "shutdown"
	METHOD_EXIT        = "exit"

	METHOD_TEXT_DOC_OPEN   = "textDocument/didOpen"
	METHOD_TEXT_DOC_CLOSE  = "textDocument/didClose"
	METHOD_TEXT_DOC_CHANGE = "textDocument/didChange"
	METHOD_TEXT_DOC_SAVE   = "textDocument/didSave"

	METHOD_TEXT_DOC_HOVER = "textDocument/hover"
)
