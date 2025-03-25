package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
