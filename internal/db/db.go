package db

import (
	"board/internal/api"
)

type DBInterface interface {
	GetUserById(id int) (api.User, error)
	GetUserByName(name string) (api.User, error)
	DeleteUserById(id int) error
	DeleteUserByName(name string) error
	AddUser(name string) error
	GetUsers() ([]api.User, error)
}
