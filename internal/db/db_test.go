package db

import (
	"testing"
	"errors"
)

type dbFactory func() DBInterface

func DBTests(t *testing.T, f dbFactory) {
	t.Helper()

	t.Run("AddUser", func (t *testing.T) { DBTestAddUser(t, f) })
	t.Run("AddUser_DuplicateFails", func (t *testing.T) { DBTestAddUser_DuplicateFails(t, f) })

	t.Run("DeleteUser", func (t *testing.T) { DBTestDeleteUser(t, f) })
	t.Run("DeleteUser_NotFound", func (t *testing.T) { DBTestDeleteUser_NotFound(t, f) })

	t.Run("GetUser", func (t *testing.T) { DBTestGetUser(t, f) })
	t.Run("GetUser_NotFound", func (t *testing.T) { DBTestGetUser_NotFound(t, f) })

	t.Run("GetUsers", func (t *testing.T) { DBTestGetUsers(t, f) })
	t.Run("GetUsers_Single", func (t *testing.T) { DBTestGetUsers_Single(t, f) })
}

func DBTestAddUser(t *testing.T, f dbFactory) {
	db := f()

	err := db.AddUser("john")

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func DBTestAddUser_DuplicateFails(t *testing.T, f dbFactory) {
	db := f()

	_ = db.AddUser("john")
	err := db.AddUser("john")

	if err == nil {
		t.Errorf("expected error for duplicate addition")
	}

	if !errors.Is(err, ErrUserExists) {
		t.Errorf("expected \"ErrUserExists\"")
	}
}

func DBTestDeleteUser(t *testing.T, f dbFactory) {
	db := f()
	_ = db.AddUser("john")

	err := db.DeleteUser("john")

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}

func DBTestDeleteUser_NotFound(t *testing.T, f dbFactory) {
	db := f()

	err := db.DeleteUser("john")

	if err != ErrUserNotFound {
		t.Errorf("expected user not found, got %v", err)
	}
}

func DBTestGetUser(t *testing.T, f dbFactory) {
	db := f()
	_ = db.AddUser("john")

	user, err := db.GetUser("john")

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if user.Name != "john" {
		t.Errorf("expected user john found %s", user.Name)
	}
}

func DBTestGetUser_NotFound(t *testing.T, f dbFactory) {
	db := f()

	_, err := db.GetUser("john")

	if err != ErrUserNotFound {
		t.Errorf("unexpected error %v", err)
	}
}

func DBTestGetUsers(t *testing.T, f dbFactory) {
	db := f()

	users, err := db.GetUsers()

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(users) != 0 {
		t.Errorf("expected no users found %v", users)
	}
}

func DBTestGetUsers_Single(t *testing.T, f dbFactory) {
	db := f()
	_ = db.AddUser("john")

	users, err := db.GetUsers()

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(users) != 1 {
		t.Errorf("expected 1 user found %v", users)
	}
}
