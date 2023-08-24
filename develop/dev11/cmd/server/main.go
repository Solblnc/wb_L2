package main

import (
	"dev11/internal"
	"log"
)

func main() {
	server, err := internal.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.SetupHandlers()
	server.Run()
}
