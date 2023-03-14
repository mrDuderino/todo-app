package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/todo-app/models"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (tip *TodoItemPostgres) CreateItem(listId int, item models.TodoItem) (int, error) {
	tx, err := tip.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	insertItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(insertItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	insertListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(insertListsItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (tip *TodoItemPostgres) GetAllItems(userId, listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done  FROM %s ti INNER JOIN %s li ON ti.id = li.item_id 
             INNER JOIN %s ul ON li.list_id = ul.list_id WHERE li.list_id = $1 AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)
	err := tip.db.Select(&items, query, listId, userId)
	return items, err
}

func (tip *TodoItemPostgres) GetById(userId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done  FROM %s ti INNER JOIN %s li ON ti.id = li.item_id 
             INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ti.id = $1 AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)
	err := tip.db.Get(&item, query, itemId, userId)
	return item, err
}

func (tip *TodoItemPostgres) Update(userId, itemId int, input models.UpdateItemInput) error {
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
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul 
                    WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ti.id = $%d AND ul.user_id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, itemId, userId)

	logrus.Debugf("update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := tip.db.Exec(query, args...)
	return err
}

func (tip *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
       WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ti.id = $1 AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)
	_, err := tip.db.Exec(query, itemId, userId)
	return err
}
