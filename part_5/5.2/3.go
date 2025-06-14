package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const tempFileName = "myTempFile.dat"
const tempDirName = "myTempDir"

func printPath(path string, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TempDir created with path: ", path)
	return path
}

func createdFile(path string) {
	filepath := filepath.Join(path, tempFileName)
	myFile, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()
	fmt.Println("TempFile created with path: ", filepath)
	fmt.Println()

}

func main() {
	os.Mkdir(tempDirName, 0777)

	path := printPath(os.MkdirTemp(tempDirName, "*")) // случайное имя временного каталога
	createdFile(path)

	path = printPath(os.MkdirTemp(tempDirName, "temp*dir")) // имя с использованием шаблона
	createdFile(path)

	path = printPath(os.MkdirTemp(tempDirName, "tempdir")) // имя с использованием шаблона
	createdFile(path)
	// создание временного каталога со случайным именем
	// в системной директории для временных файлов
	path = printPath(os.MkdirTemp("", "*"))
	createdFile(path)
}
