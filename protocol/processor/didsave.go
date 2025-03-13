package processor

import (
	"github.com/sskender/jargonlsp/protocol/base"
	"github.com/sskender/jargonlsp/state"
)

type DidSaveTextDocumentParams struct {
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
	Text         string                  `json:"text"`
}

type DidSaveTextDocumentNotification struct {
	base.NotificationMessage
	Params *DidSaveTextDocumentParams `json:"params"`
}

func DocumentDidSave(notificationMessage any) (any, error) {
	notification := notificationMessage.(*DidSaveTextDocumentNotification)

	key := notification.Params.TextDocument.Uri
	content := notification.Params.Text

	gdb := state.GetDatabase()

	err := gdb.Documents.Update(key, content, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
