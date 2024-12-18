package state

import (
	"fmt"
	"log"
	"sync"
)

// TODO implement locks on all api calls

var lock = &sync.Mutex{}

var globalState *StateDB

func GetState() *StateDB {
	if globalState == nil {

		lock.Lock()
		defer lock.Unlock()

		if globalState == nil {
			log.Println("initializing global state for the first time")

			globalState = &StateDB{
				Documents: map[string]*DocumentItem{},
			}
		}
	}

	return globalState
}

func (s *StateDB) Count() int {
	return len(s.Documents)
}

func (s *StateDB) IsEmpty() bool {
	return s.Count() == 0
}

func (s *StateDB) Save(key string, doc *DocumentItem) error {
	if s.Documents[key] != nil {
		return fmt.Errorf("save failed: already exists %s", key)
	}

	s.Documents[key] = doc

	log.Println("saved content of document", key)
	log.Printf("global state is now managing %d files", globalState.Count())

	return nil
}

func (s *StateDB) Update(key string, content string, version int) error {
	if s.Documents[key] == nil {
		return fmt.Errorf("update failed: invalid key %s", key)
	}

	if s.Documents[key].Version >= version {
		return fmt.Errorf("update failed: old version")
	}

	log.Printf("updating content of document %s with version %d", key, s.Documents[key].Version)

	s.Documents[key].Version = version
	s.Documents[key].Text = content

	log.Printf("updated content of document %s to version %d", key, s.Documents[key].Version)

	return nil
}
