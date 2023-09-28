package main

import (
	"log"
	"os"
)

func main() {
	defer log.Println("Done")

	args := parseArguments()
	initializeLogging(args.logTimestamps)
	log.Println("Starting")
}

func initializeLogging(logTimestamps bool) {
	if logTimestamps {
		log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.Lshortfile)
	}
	log.SetOutput(os.Stdout)
}
