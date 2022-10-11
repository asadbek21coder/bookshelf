package main

import (
	"log"

	"github.com/asadbek21coder/bookshelf"
	"github.com/asadbek21coder/bookshelf/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(bookshelf.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
