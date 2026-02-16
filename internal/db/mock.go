package db

import (
	"board/api"
	"slices"
)

type DBMock struct {
	user_id int
	users   []api.User
}

func NewDBMock() DBMock {
	return DBMock {
		user_id : 0,
	}
}

func (db *DBMock) AddUser(name string) error {
	_, err := db.GetUser(name)
	if err == nil {
		return ErrUserExists
	}

	id := db.user_id
	db.user_id += 1
	user := api.User { Id : id, Name : name }
	db.users = append(db.users, user)
	return nil
}

func (db *DBMock) GetUser(name string) (api.User, error) {
	idx := slices.IndexFunc(db.users, func (u api.User) bool { return u.Name == name })
	if idx == -1 {
		return api.User {}, ErrUserNotFound
	}
	return db.users[idx], nil
}

func (db *DBMock) DeleteUser(name string) error {
	idx := slices.IndexFunc(db.users, func (u api.User) bool { return u.Name == name })
	if idx == -1 {
		return ErrUserNotFound
	}
	db.users = append(db.users[:idx], db.users[idx+1:]...)
	return nil
}

func (db *DBMock) GetUsers() ([]api.User, error) {
	return db.users, nil
}
