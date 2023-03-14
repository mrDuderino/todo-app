package models

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int `db:"user_id"`
	ListId int `db:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (uii *UpdateItemInput) Validate() error {
	if uii.Title == nil && uii.Description == nil && uii.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type ListsItem struct {
	Id     int `db:"id"`
	ListId int `db:"list_id"`
	ItemId int `db:"item_id"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (uli *UpdateListInput) Validate() error {
	if uli.Title == nil && uli.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
