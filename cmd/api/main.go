package main

import (
	"log"

	"github.com/alsolovyev/dummy-api/internal/transport/rest"
	"github.com/alsolovyev/dummy-api/internal/transport/rest/handler"
)

const (
	address = "127.0.0.1"
	port    = 8080
)

func main() {
	handler := new(handler.Handler)

	if err := new(rest.Server).Run(address, port, handler.MakeRouter()); err != nil {
		log.Fatalf("Error occured while running HTTP server: %s", err.Error())
	}
}
