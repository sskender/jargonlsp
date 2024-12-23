package processor

import (
	"jargonlsp/protocol/base"
	"log"
)

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

const MARKUP_KIND = "markdown"

const MARKUP_VALUE = `
# This is a title

*This is just a longer paragraph that gives the definition for hovered text.*
`

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

	// TODO implement

	key := message.Params.TextDocument.Uri

	log.Println(key)

	line := message.Params.Position.Line
	character := message.Params.Position.Character

	result := HoverResponse{
		Contents: &MarkupContent{
			Kind:  MARKUP_KIND,
			Value: MARKUP_VALUE,
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
