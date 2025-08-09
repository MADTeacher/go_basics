package main

import (
	"log"
	"os"
)

func main() {
	pathToFile := "test.txt"
	err := os.Remove(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
}
