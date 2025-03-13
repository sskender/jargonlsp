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

type Documents map[string]*DocumentItem

func (d Documents) Count() int {

	if d == nil {
		return 0
	}

	return len(d)
}

func (d Documents) Get(key string) (*DocumentItem, error) {

	if d == nil {
		return nil, fmt.Errorf("get failed: documents are not initialized")
	}

	item, exists := d[key]
	if !exists {
		return nil, fmt.Errorf("get failed: invalid key %s", key)
	}

	return item, nil
}

func (d Documents) Save(key string, doc *DocumentItem) error {

	if d == nil {
		return fmt.Errorf("save failed: documents are not initialized")
	}

	_, exists := d[key]

	if exists {
		return fmt.Errorf("save failed: already exists %s", key)
	}

	d[key] = doc

	log.Printf("document saved %s", key)
	log.Printf("database is managing %d files", d.Count())

	return nil
}

func (d Documents) Delete(key string) error {

	if d == nil {
		return fmt.Errorf("delete failed: documents are not initialized")
	}

	_, exists := d[key]
	if !exists {
		return fmt.Errorf("delete failed: invalid key %s", key)
	}

	log.Printf("database is managing %d files", d.Count())

	// TODO delete is not correct
	delete(d, key)

	log.Printf("document deleted %s", key)
	log.Printf("database is managing %d files", d.Count())

	return nil
}

func (d Documents) Update(key string, content string, version *uint) error {

	if d == nil {
		return fmt.Errorf("update failed: documents are not initialized")
	}

	item, exists := d[key]
	if !exists {
		return fmt.Errorf("update failed: invalid key %s", key)
	}

	if version != nil && item.Version >= *version {
		return fmt.Errorf("update failed: version too old")
	}

	log.Printf("updating content of document %s with version %d", key, item.Version)

	item.Text = content

	if version == nil {
		item.Version++
	} else {
		item.Version = *version
	}

	log.Printf("updated content of document %s to version %d", key, item.Version)

	return nil
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
		log.Println("no token selected")
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
