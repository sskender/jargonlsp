package processor

import (
	"jargonlsp/protocol/base"
	"jargonlsp/state"
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

	gstate := state.GetState()

	err := gstate.Update(key, content, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
