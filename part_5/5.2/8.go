package main

import (
	"fmt"
	"log"
	"os"
)

func getWD() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current path:", path)
}

func main() {
	getWD()
	os.Mkdir("myDir", 0777)
	os.Chdir("./myDir")
	getWD()
	myFile, err := os.Create("test.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()
}
