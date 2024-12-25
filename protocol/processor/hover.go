package processor

import (
	"fmt"
	"jargonlsp/protocol/base"
	"jargonlsp/state"
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

	gstate := state.GetState()

	token, err := gstate.GetToken(key, line, character)
	if err != nil {
		return nil, err
	}

	// TODO perform lookup and return definition

	result := HoverResponse{
		Contents: &MarkupContent{
			Kind:  MARKUP_KIND,
			Value: formatMarkup(*token),
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

func formatMarkup(token string) string {
	return fmt.Sprintf(`
# Hello Hover

*The selected token is '%s'.*
`, token)
}
