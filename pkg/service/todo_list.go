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

func (tls *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return tls.repo.GetAll(userId)
}

func (tls *TodoListService) GetById(userId, listId int) (models.TodoList, error) {
	return tls.repo.GetById(userId, listId)
}
