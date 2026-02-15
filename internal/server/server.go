package server

import (
	"board/internal/db"
)

type Server struct{
	db db.DBInterface
}

func ServerNew(db db.DBInterface) Server {
	return Server {
		db: db,
	}
}
