package state

import (
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type Database struct {
	Documents  Documents
	Dictionary Dictionary
}

// TODO implement locks on all calls

var gdb *Database

func GetDatabase() *Database {
	if gdb == nil {

		lock.Lock()
		defer lock.Unlock()

		if gdb == nil {
			log.Println("initializing database for the first time")

			gdb = &Database{
				Documents:  Documents{},
				Dictionary: Dictionary{},
			}
		}
	}

	return gdb
}
