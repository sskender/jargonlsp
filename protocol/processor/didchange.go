package processor

import (
	"github.com/sskender/jargonlsp/protocol/base"
	"github.com/sskender/jargonlsp/state"
)

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}

type TextDocumentIdentifier struct {
	Uri string `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version uint `json:"version"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type DidChangeTextDocumentNotification struct {
	base.NotificationMessage
	Params *DidChangeTextDocumentParams `json:"params"`
}

func DocumentDidChange(notificationMessage any) (any, error) {
	notification := notificationMessage.(*DidChangeTextDocumentNotification)

	key := notification.Params.TextDocument.Uri
	version := notification.Params.TextDocument.Version

	gdb := state.GetDatabase()

	for _, item := range notification.Params.ContentChanges {
		err := gdb.Documents.Update(key, item.Text, &version)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
