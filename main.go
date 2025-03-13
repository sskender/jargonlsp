package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sskender/jargonlsp/server"
)

const (
	LOG_FILE_PATH = "debug.log"
)

func main() {

	// TODO logging is bad

	logFile, err := os.OpenFile(LOG_FILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

	// parse flags

	showVersion := flag.Bool("version", false, "Show version")
	dictionaryPath := flag.String("dictionary", "", "Dictionary file")

	flag.Parse()

	// server settings

	settings := server.ServerSettings{
		DictionaryPath: dictionaryPath,
		EnableTcp:      false,
	}

	// init server

	jargon := server.New(settings)

	if *showVersion {
		fmt.Println(jargon.Version())
		return
	}

	// run main lopp

	jargon.RunLoop()

	// TODO proper exit from all routines

	log.Println("exiting gracefully")
}
