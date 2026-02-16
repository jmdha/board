package db

import (
	"testing"
)

func dbMockFactory() DBInterface {
	db := DBMock {}
	return &db
}

func TestMockInterface(t *testing.T) {
	t.Run("DBMock", func (t *testing.T) { DBTests(t, dbMockFactory) })
}
