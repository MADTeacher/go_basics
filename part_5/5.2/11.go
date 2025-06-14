package main

import (
	"log"
	"os"
)

func createDirWithFile() {
	os.Mkdir("myDir", 0777)
	myFile, err := os.Create("myDir\\test.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()
}

func main() {
	createDirWithFile()
	os.Rename("myDir", "newDir") // переименование
	createDirWithFile()
	os.Rename("myDir", "newDir\\myDir") // перемещение
}
