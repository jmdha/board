package server

import (
	"database/sql"
)

type Server struct{
	db *sql.DB
}

func ServerNew(db_path string) (Server, error) {
	db, err := db_conn(db_path)
	if err != nil {
		return Server {}, err
	}

	return Server {
		db: db,
	}, nil
}
