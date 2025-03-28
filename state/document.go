package state

import (
	"fmt"
	"log"
	"strings"
)

type DocumentItem struct {
	Key        string
	LanguageId string
	Text       string
	Version    uint
}

func (d *DocumentItem) GetToken(line uint, column uint) (*string, error) {

	lines := strings.Split(d.Text, "\n")

	if len(lines) < int(line) {
		return nil, fmt.Errorf("invalid line %d for key %s", line, d.Key)
	}

	textLine := lines[line]

	if len(textLine) < int(column) {
		return nil, fmt.Errorf("invalid column %d for key %s", column, d.Key)
	}

	token := getTokenFromText(textLine, column)

	if token == nil {
		log.Println("no token is selected")
	} else {
		log.Printf("selected token is '%s'", *token)
	}

	return token, nil
}

func getTokenFromText(textLine string, cursor uint) *string {

	if !isPartOfToken(rune((textLine)[cursor])) {
		return nil
	}

	colStart, colEnd := cursor, cursor

	for colStart > 0 && isPartOfToken(rune((textLine)[colStart-1])) {
		colStart--
	}

	for colEnd < uint(len(textLine)) && isPartOfToken(rune((textLine)[colEnd])) {
		colEnd++
	}

	if colStart == colEnd {
		return nil
	}

	token := (textLine)[colStart:colEnd]

	return &token
}

func isPartOfToken(c rune) bool {
	return (c >= 97 && c <= 122) || (c >= 65 && c <= 90) || (c >= 48 && c <= 57) || c == 95
}
