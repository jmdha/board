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
}

func DBTestAddUser(t *testing.T, f dbFactory) {
	db := f()

	err := db.AddUser("john")

	if err != nil {
		t.Errorf("unexpected AddUser error %v", err)
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
