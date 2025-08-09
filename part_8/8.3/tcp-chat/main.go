package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tcp/server/client"
	"tcp/server/server"
)

const PORT = 8081

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You didn't select a launch option!!!")
		return
	}
	numGR, err := strconv.Atoi(os.Args[1])
	checkErr(err)

	switch numGR {
	case 1:
		fmt.Println("Server is running")
		server.RunServer(PORT)
	case 2:
		fmt.Println("Client is running")
		client.RunClient(PORT)
	default:
		log.Fatal("What pokemon is this?")
	}
}
