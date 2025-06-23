package main

import (
	"flag"
	"log"

	"tic-tac-toe/client"
	"tic-tac-toe/database"
	"tic-tac-toe/server"
)

func main() {
	// Флаг для выбора режима работы
	mode := flag.String(
		"mode", "server",
		"start in 'server' or 'client' mode",
	)
	// Флаг для указания адреса
	// :8088 = 127.0.0.1:8088 или localhost:8088
	addr := flag.String("addr", ":8088", "address to run on")
	flag.Parse() // Парсим флаги

	switch *mode { // Переключение режима работы
	case "server": // Режим сервера
		// Создаем репозиторий
		repository, err := database.NewSQLiteRepository()
		if err != nil {
			log.Fatalf("Failed to create repository: %v", err)
		}
		// Создаем сервер и запускаем его, передав
		// ему адрес в формате "ip:port" и репозиторий
		srv, err := server.NewServer(*addr, repository)
		if err != nil {
			log.Fatalf("Failed to create server: %v", err)
		}
		srv.Start()
	case "client": // Режим клиента
		// Создаем клиента и запускаем его, передав
		// ему адрес сервера в формате "ip:port"
		cli, err := client.NewClient(*addr)
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		cli.Start()
	default: // Неизвестный режим
		log.Fatalf(
			"Unknown mode: %s. Use 'server' or 'client'.", *mode,
		)
	}
}
