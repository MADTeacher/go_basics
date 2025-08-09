package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	pathToFile := "pirates.txt"
	fileInfo, err := os.Stat(pathToFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
			os.Exit(1) // выход  из приложения со статусом 1
		}
	}
	fmt.Println("File does exist!")
	fmt.Println("File information:", fileInfo)
}
