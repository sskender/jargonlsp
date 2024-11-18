package main

import (
	"jargonlsp/server"
	"log"
	"os"
)

const (
	LOG_FILE_PATH = "debug.log"
)

func main() {
	logFile, err := os.OpenFile(LOG_FILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

	log.Println("JargonLSP main process started")

	jargon := server.New()
	jargon.Run()

	log.Println("NEVER REACHED")
}
