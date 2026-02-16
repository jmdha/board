package db

import (
	"log"
	"testing"
)

func dbSqliteFactory() DBInterface {
	db := DBSqlite {}
	err := db.Create(":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &db
}

func TestSqliteInterface(t *testing.T) {
	t.Run("DBSqlite", func (t *testing.T) { DBTests(t, dbSqliteFactory) })
}
