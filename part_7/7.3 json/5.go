package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name        string   `json:"name"` // имя поля при сериализации/десериализации
	Age         uint8    `json:"-"`    // поле не должно учитываться
	Course      uint8    `json:"-"`
	Single      bool     `json:"-"`
	Description []string `json:"description"`
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

	data, err := json.MarshalIndent(student, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
