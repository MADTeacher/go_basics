package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("pirates2.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
