package main

import (
	"encoding/json"
	"log"
	"os"
)

type Student struct {
	Name        string
	Age         uint8
	Course      uint8
	Single      bool
	Description []string
}

func main() {
	student := &Student{
		Name:   "Alex",
		Age:    20,
		Course: 2,
		Single: true,
		Description: []string{
			"Мечтатель",
			"Ленив",
			"Студент",
			"Постоянно жалуется на жизнь",
		},
	}

	data, err := json.Marshal(student)
	myFile, err := os.Create("student.json")
	if err != nil {
		log.Fatal(err)
	}
	myFile.Write(data)
	defer myFile.Close()
}
