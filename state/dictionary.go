package state

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Dictionary map[string]string

func (d Dictionary) Size() int {
	if d == nil {
		return 0
	}

	return len(d)
}

func (d Dictionary) Load(filepath *string) error {

	if filepath == nil || len(*filepath) == 0 {
		// TODO return error or not ???
		log.Println("warning: dictionary is not defined")
		return nil
	}

	log.Printf("loading dictionary from %s", *filepath)

	jsonFile, err := os.Open(*filepath)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&d)
	if err != nil {
		return err
	}

	log.Printf("loaded %d words into dictionary from %s", d.Size(), *filepath)

	return nil
}

func (d Dictionary) GetDefinition(lexem *string) (*string, error) {

	if d == nil {
		return nil, errors.New("get definition failed: dictionary is not initialized")
	}

	if lexem == nil {
		return nil, nil
	}

	if definition, found := d[*lexem]; found {
		return &definition, nil
	}

	return nil, nil
}
