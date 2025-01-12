package state

import (
	"fmt"
	"log"
	"sync"
)

type StateDB struct {
	Documents map[string]*DocumentItem
}

// TODO implement locks on all api calls

var lock = &sync.Mutex{}

var globalState *StateDB

func GetState() *StateDB {
	if globalState == nil {

		lock.Lock()
		defer lock.Unlock()

		if globalState == nil {
			log.Println("initializing database for the first time")

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

func (s *StateDB) Get(key string) (*DocumentItem, error) {
	if s.Documents[key] == nil {
		return nil, fmt.Errorf("get failed: invalid key %s", key)
	}

	return s.Documents[key], nil
}

func (s *StateDB) Save(key string, doc *DocumentItem) error {
	if s.Documents[key] != nil {
		return fmt.Errorf("save failed: already exists %s", key)
	}

	s.Documents[key] = doc

	log.Printf("document saved %s", key)
	log.Printf("database is managing %d files", s.Count())

	return nil
}

func (s *StateDB) Delete(key string) error {
	if s.Documents[key] == nil {
		return fmt.Errorf("delete failed: invalid key %s", key)
	}

	log.Printf("database is managing %d files", s.Count())

	// TODO does not delete just update version

	delete(s.Documents, key)

	log.Printf("document deleted %s", key)
	log.Printf("database is managing %d files", s.Count())

	return nil
}

func (s *StateDB) Update(key string, content string, version *uint) error {
	if s.Documents[key] == nil {
		return fmt.Errorf("update failed: invalid key %s", key)
	}

	if version != nil && s.Documents[key].Version >= *version {
		return fmt.Errorf("update failed: version too old")
	}

	log.Printf("updating content of document %s with version %d", key, s.Documents[key].Version)

	s.Documents[key].Text = content

	if version == nil {
		s.Documents[key].Version++
	} else {
		s.Documents[key].Version = *version
	}

	log.Printf("updated content of document %s to version %d", key, s.Documents[key].Version)

	return nil
}
