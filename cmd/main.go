package main

import (
	"flag"
	"log"
	"fmt"
	"net/http"
	"board/internal/api"
	"board/internal/db"
	"board/internal/server"
)

func main() {
	var port int
	var db_path string

	flag.IntVar(&port, "p", 8080, "which port to operate on")
	flag.StringVar(&db_path, "d", "db.sqlite", "path to db")
	flag.Parse()

	db := &db.DBSqlite {}
	log.Fatal(db.Create(db_path))
	log.Fatal(db.Connect(db_path))

	s := server.ServerNew(db)
	handler := api.Handler(&s)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
