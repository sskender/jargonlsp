package processor

import (
	"jargonlsp/protocol/base"
	"jargonlsp/state"
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

	gstate := state.GetState()

	for _, item := range notification.Params.ContentChanges {
		err := gstate.Update(key, item.Text, &version)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
