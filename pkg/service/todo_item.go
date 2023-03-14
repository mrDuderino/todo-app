package service

import (
	"github.com/mrDuderino/todo-app/models"
	"github.com/mrDuderino/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (tis *TodoItemService) CreateItem(userId int, listId int, item models.TodoItem) (int, error) {
	_, err := tis.listRepo.GetById(userId, listId)
	if err != nil {
		//list does not exist ordoes notbelongs to user
		return 0, err
	}
	return tis.repo.CreateItem(listId, item)
}

func (tis *TodoItemService) GetAllItems(userId int, listId int) ([]models.TodoItem, error) {
	return tis.repo.GetAllItems(userId, listId)
}

func (tis *TodoItemService) GetById(userId, itemId int) (models.TodoItem, error) {
	return tis.repo.GetById(userId, itemId)
}
