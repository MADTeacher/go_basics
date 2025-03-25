package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const tempFileName = "myTempFile.dat"

func createdFile(path string) {
	filepath := filepath.Join(path, tempFileName)
	myFile, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()
	fmt.Println("File created with path: ", filepath)
}

func main() {
	path := filepath.Join("firstDir", "secondDir", "new")
	os.MkdirAll(path, 0777)
	createdFile(path)
}
