package main

import (
	"log"

	"github.com/alsolovyev/dummy-api/internal/transport/rest"
)

const (
	address = "127.0.0.1"
	port    = 8080
)

func main() {
	if err := new(rest.Server).Run(address, port); err != nil {
		log.Fatalf("Error occured while running HTTP server: %s", err.Error())
	}
}
