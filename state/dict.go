package state

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Dictionary map[string]string

func (d *Dictionary) Size() int {
	return len(*d)
}

func (d *Dictionary) Load(filepath *string) error {

	if filepath == nil || len(*filepath) == 0 {
		fmt.Println("warning: dictionary is not defined")
		return nil
	}

	log.Printf("loading dictionary from %s", *filepath)

	jsonFile, err := os.Open(*filepath)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(d)
	if err != nil {
		return err
	}

	log.Printf("loaded %d words into dictionary from %s", d.Size(), *filepath)

	return nil
}

func (d *Dictionary) GetDefinition(lexem *string) (*string, error) {

	if lexem == nil {
		return nil, nil
	}

	if definition, found := (*d)[*lexem]; found {
		return &definition, nil
	}

	return nil, nil
}

var globalDictionary *Dictionary

func GetDictionary() *Dictionary {
	if globalDictionary == nil {

		lock.Lock()
		defer lock.Unlock()

		if globalDictionary == nil {
			log.Println("initializing dictionary for the first time")

			globalDictionary = &Dictionary{}
		}
	}

	return globalDictionary
}
