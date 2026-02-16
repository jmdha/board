package db

import (
	"errors"
	"board/api"
)

type DBInterface interface {
	GetUserById(id int) (api.User, error)
	GetUserByName(name string) (api.User, error)
	DeleteUserById(id int) error
	DeleteUserByName(name string) error
	AddUser(name string) error
	GetUsers() ([]api.User, error)
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserExists = errors.New("user already exists")
