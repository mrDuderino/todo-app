package service

import (
	"github.com/mrDuderino/todo-app/models"
	"github.com/mrDuderino/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (int, error)
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
