package main

import (
	"golang/todo-service/db"
	"golang/todo-service/service"
	"log"
	"net/http"
)

const PORT = "8080"

func main() {

	rep := db.NewSQLiteRepository()
	defer rep.Close()

	router := service.NewRouter(rep)
	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
