package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	file.Close() // закрытие файла
}
