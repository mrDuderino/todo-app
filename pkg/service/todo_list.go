package service

import (
	"github.com/mrDuderino/todo-app/models"
	"github.com/mrDuderino/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (tls *TodoListService) CreateList(userId int, list models.TodoList) (int, error) {
	return tls.repo.CreateList(userId, list)
}
