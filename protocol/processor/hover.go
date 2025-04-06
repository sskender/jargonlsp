package processor

import (
	"fmt"

	"github.com/sskender/jargonlsp/protocol/base"
	"github.com/sskender/jargonlsp/state"
)

const MARKUP_KIND = "markdown"

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

type Range struct {
	Start *Position `json:"start"`
	End   *Position `json:"end"`
}

type TextDocumentPositionParams struct {
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
	Position     *Position               `json:"position"`
}

type HoverParams struct {
	TextDocumentPositionParams
}

type HoverRequest struct {
	base.RequestMessage
	Params *HoverParams `json:"params"`
}

type MarkupContent struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

type HoverResponse struct {
	Contents *MarkupContent `json:"contents"`
	Range    *Range         `json:"range"`
}

func DocumentHover(requestMessage any) (any, error) {
	message := requestMessage.(*HoverRequest)

	key := message.Params.TextDocument.Uri

	line := message.Params.Position.Line
	character := message.Params.Position.Character

	gdb := state.GetDatabase()

	document, err := gdb.Documents.Get(key)
	if err != nil {
		return nil, err
	}

	token, err := document.GetToken(line, character)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, nil
	}

	definition, err := gdb.Dictionary.GetDefinition(*token)
	if err != nil {
		return nil, err
	}

	if definition == nil {
		return nil, nil
	}

	result := HoverResponse{
		Contents: &MarkupContent{
			Kind:  MARKUP_KIND,
			Value: formatMarkupResponse(*token, *definition),
		},
		Range: &Range{
			Start: &Position{
				Line:      line,
				Character: character - 1,
			},
			End: &Position{
				Line:      line,
				Character: character + 1,
			},
		},
	}

	return result, nil
}

// TODO markup looks buggy

func formatMarkupResponse(token string, definition string) string {
	return fmt.Sprintf(`
# *%s*

%s
`, token, definition)
}
