package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("pirates.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("jgjgjg") // ничего в файл не добавится!!!
	data := make([]byte, 128)
	for {
		length, err := file.Read(data)
		fmt.Printf("Reading %d byte\n", length)
		if err == io.EOF { // достигли конца файла?
			break
		}
	}
	fmt.Println(string(data))

	fmt.Println("Text was read from file")
}
