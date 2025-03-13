package processor

import (
	"github.com/sskender/jargonlsp/protocol/base"
	"github.com/sskender/jargonlsp/state"
)

type DidCloseTextDocumentParams struct {
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

type DidCloseTextDocumentNotification struct {
	base.NotificationMessage
	Params *DidCloseTextDocumentParams `json:"params"`
}

func DocumentDidClose(notificationMessage any) (any, error) {
	notification := notificationMessage.(*DidCloseTextDocumentNotification)

	key := notification.Params.TextDocument.Uri

	gdb := state.GetDatabase()

	err := gdb.Documents.Delete(key)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
