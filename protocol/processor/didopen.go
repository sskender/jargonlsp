package processor

import (
	"github.com/sskender/jargonlsp/protocol/base"
	"github.com/sskender/jargonlsp/state"
)

type TextDocumentItem struct {
	Uri        string `json:"uri"`
	LanguageId string `json:"languageId"`
	Version    uint   `json:"version"`
	Text       string `json:"text"`
}

type DidOpenTextDocumentParams struct {
	TextDocument *TextDocumentItem `json:"textDocument"`
}

type DidOpenTextDocumentNotification struct {
	base.NotificationMessage
	Params *DidOpenTextDocumentParams `json:"params"`
}

func DocumentDidOpen(notificationMessage any) (any, error) {
	notification := notificationMessage.(*DidOpenTextDocumentNotification)

	key := notification.Params.TextDocument.Uri

	doc := state.DocumentItem{
		Key:        key,
		LanguageId: notification.Params.TextDocument.LanguageId,
		Text:       notification.Params.TextDocument.Text,
		Version:    notification.Params.TextDocument.Version,
	}

	gdb := state.GetDatabase()

	err := gdb.Documents.Save(key, &doc)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
