package repository

import (
	"fmt"
	"github.com/mrDuderino/todo-app/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES($1, $2, $3) RETURNING id", usersTable)
	row := ap.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ap *AuthPostgres) GetUser(userName, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", usersTable)
	err := ap.db.Get(&user, query, userName, password)
	return user, err
}
