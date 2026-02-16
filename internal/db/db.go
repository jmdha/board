package db

import (
	"errors"
	"board/api"
)

type DBInterface interface {
	AddUser(name string) error
	DeleteUser(name string) error
	GetUser(name string) (api.User, error)
	GetUsers() ([]api.User, error)
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserExists = errors.New("user already exists")
