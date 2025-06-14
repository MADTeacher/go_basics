package main

import (
	"fmt"
	"log"
	"os"
)

func getWD() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current path:", path)
}

func main() {
	getWD()
	os.Chdir("../")
	getWD()
}
