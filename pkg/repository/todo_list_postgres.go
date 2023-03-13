package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/todo-app/models"
	"github.com/sirupsen/logrus"
	"strings"
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

func (tlp *TodoListPostgres) Update(userId int, listId int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := tlp.db.Exec(query, args...)
	return err
}

func (tlp *TodoListPostgres) Delete(userId, listId int) error {
	deleteListQuery := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND tl.id = $2", todoListsTable, usersListsTable)
	_, err := tlp.db.Exec(deleteListQuery, userId, listId)

	return err
}
