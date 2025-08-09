package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	originalFile, err := os.Open("pirates2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("pirates2_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copied %d bytes", bytesWritten) // Copied 224 bytes

	// Фиксируем содержимое файла и осуществляем освобождение буфера
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
