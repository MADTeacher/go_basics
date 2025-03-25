package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("pirates.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
