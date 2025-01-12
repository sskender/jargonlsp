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
	logFile, err := os.OpenFile(LOG_FILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

	jargon := server.New()

	showVersion := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *showVersion {
		fmt.Println(jargon.Version())
		os.Exit(0)
	}

	jargon.Run()

	log.Println("exiting gracefully")
}
