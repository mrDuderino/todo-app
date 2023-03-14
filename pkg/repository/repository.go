package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/todo-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(userName, password string) (models.User, error)
}

type TodoList interface {
	CreateList(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Update(userId int, listId int, input models.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	CreateItem(listId int, item models.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
