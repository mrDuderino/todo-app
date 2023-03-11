package main

import (
	"log"

	"github.com/mrDuderino/todo-app"
	"github.com/mrDuderino/todo-app/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
