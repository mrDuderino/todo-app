package main

import (
	"log"

	"github.com/mrDuderino/todo-app"
	"github.com/mrDuderino/todo-app/pkg/handler"
	"github.com/mrDuderino/todo-app/pkg/repository"
	"github.com/mrDuderino/todo-app/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
