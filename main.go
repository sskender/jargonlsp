package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sskender/jargonlsp/server"
	"github.com/sskender/jargonlsp/state"
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

	// TODO logging is bad to here

	showVersion := flag.Bool("version", false, "Show version")
	dictionaryPath := flag.String("dict", "", "Dictionary file to load")

	// TODO support flag --stdio

	flag.Parse()

	jargon := server.New()

	if *showVersion {
		fmt.Println(jargon.Version())
		os.Exit(0)
	}

	err = state.GetDictionary().Load(dictionaryPath)
	if err != nil {
		panic(err)
	}

	jargon.RunLoop()

	// TODO proper exit from all routines

	log.Println("exiting gracefully")
}
