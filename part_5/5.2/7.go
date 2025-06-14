package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path) // e:\code\golang\directory
}
