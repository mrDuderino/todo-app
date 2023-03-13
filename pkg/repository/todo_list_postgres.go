package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/todo-app/models"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (tlp *TodoListPostgres) CreateList(userId int, list models.TodoList) (int, error) {
	tx, err := tlp.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (tlp *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList
	todoListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)
	err := tlp.db.Select(&lists, todoListsQuery, userId)
	return lists, err
}

func (tlp *TodoListPostgres) GetById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList
	todoListQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2", todoListsTable, usersListsTable)
	err := tlp.db.Get(&list, todoListQuery, userId, listId)
	return list, err
}
