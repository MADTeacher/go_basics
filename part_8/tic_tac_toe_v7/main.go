package main

import (
	"flag"
	"log"

	"tic-tac-toe/client"
	"tic-tac-toe/database"
	"tic-tac-toe/server"
)

func main() {
	mode := flag.String(
		"mode", "server",
		"start in 'server' or 'client' mode",
	)
	addr := flag.String("addr", ":8088", "address to run on")
	flag.Parse()

	switch *mode {
	case "server":
		repository, err := database.NewSQLiteRepository()
		if err != nil {
			log.Fatalf("Failed to create repository: %v", err)
		}
		srv, err := server.NewServer(*addr, repository)
		if err != nil {
			log.Fatalf("Failed to create server: %v", err)
		}
		srv.Start()
	case "client":
		cli, err := client.NewClient(*addr)
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		cli.Start()
	default:
		log.Fatalf(
			"Unknown mode: %s. Use 'server' or 'client'.", *mode,
		)
	}
}
