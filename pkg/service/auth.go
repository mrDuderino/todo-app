package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/mrDuderino/todo-app"
	"github.com/mrDuderino/todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = as.generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv("SALT")
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
