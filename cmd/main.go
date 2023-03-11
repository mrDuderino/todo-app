package main

import (
	"log"

	"github.com/mrDuderino/todo-app"
	"github.com/mrDuderino/todo-app/pkg/handler"
	"github.com/mrDuderino/todo-app/pkg/repository"
	"github.com/mrDuderino/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
