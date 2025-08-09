package main

import (
	"golang/todo/db"
	"golang/todo/menu"
	"time"
)

func main() {
	rep := db.NewSQLiteRepository()
	defer rep.Close()
	for {
		menu.CreateMenu(rep)
		time.Sleep(2 * time.Second)
	}
}
