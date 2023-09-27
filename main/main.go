package main

import "log"

func main() {
	defer log.Println("Done")

	log.Println("Starting")
}
