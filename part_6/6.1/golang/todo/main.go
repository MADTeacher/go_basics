package main

import (
	"golang/todo/db"
	"golang/todo/menu"
	"time"
)

func main() {
	// Создаем репозиторий
	rep := db.NewSQLiteRepository()
	// Создаем отложенное закрытие соединения
	defer rep.Close()
	// Бесконечный цикл
	for {
		menu.CreateMenu(rep)
		time.Sleep(2 * time.Second)
	}
}
