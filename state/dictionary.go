package state

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

type Dictionary map[string]string

func (d Dictionary) Size() int {
	if d == nil {
		return 0
	}

	return len(d)
}

func (d Dictionary) Load(filepath string) error {

	log.Printf("loading dictionary from '%s'", filepath)

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	tmpDict := Dictionary{}

	err = json.NewDecoder(jsonFile).Decode(&tmpDict)
	if err != nil {
		return err
	}

	for k, v := range tmpDict {
		d[strings.ToLower(k)] = v
	}

	log.Printf("loaded %d words into dictionary from '%s'", d.Size(), filepath)

	return nil
}

func (d Dictionary) GetDefinition(lexem string) (*string, error) {

	if d == nil {
		return nil, errors.New("get definition failed: dictionary is not initialized")
	}

	lexem = strings.ToLower(lexem)

	if definition, found := d[lexem]; found {
		return &definition, nil
	}

	return nil, nil
}
