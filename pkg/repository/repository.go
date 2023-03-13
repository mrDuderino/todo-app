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
}

type TodoItem interface {
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
	}
}
