package main

import (
	"flag"
	"log"
	"fmt"
	"net/http"
	"board/internal/api"
	"board/internal/server"
)

func main() {
	var port int
	var db_path string

	flag.IntVar(&port, "p", 8080, "which port to operate on")
	flag.StringVar(&db_path, "d", "db.sqlite", "path to db")
	flag.Parse()

	s, err := server.ServerNew(db_path)
	if err != nil {
		log.Fatal("failed to create server with error " + err.Error())
	}
	handler := api.Handler(&s)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
