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

	tempData := make([]byte, 32)
	data := []byte{}
	for {
		length, err := file.Read(tempData)
		fmt.Printf("Reading %d byte\n", length)
		if err == io.EOF { // достигли конца файла?
			break
		}
		data = append(data, tempData...)
	}
	fmt.Println(string(data))
}
