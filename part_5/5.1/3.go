package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test.txt"
	newPath := "../test2.txt" // переместит файл на уровень выше с именем test2.txt
	// можно также оставить старое имя у перемещаемого файла - test.txt
	// newPath := "test2.txt" // переименование файла в текущей директории
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
