package main

import (
	"log"

	"github.com/alsolovyev/dummy-api/internal/repository/sqlite"
	"github.com/alsolovyev/dummy-api/internal/transport/rest"
	"github.com/alsolovyev/dummy-api/internal/transport/rest/handler"
)

const (
	address = "127.0.0.1"
	port    = 8080
)

func main() {
	_, err := sqlite.MakeRepository(sqlite.Config{
		Path: "./.bin/sqlite.db",
	})

	if err != nil {
		log.Fatalf("Error occurred while creating repository: %s", err.Error())
	}

	handler := new(handler.Handler)

	if err := new(rest.Server).Run(address, port, handler.MakeRouter()); err != nil {
		log.Fatalf("Error occurred while running HTTP server: %s", err.Error())
	}
}
