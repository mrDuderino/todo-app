package service

import (
	"github.com/mrDuderino/todo-app"
	"github.com/mrDuderino/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service { // dependency injection example
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
