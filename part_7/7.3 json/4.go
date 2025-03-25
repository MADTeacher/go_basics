package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name        string   `json:"name"` // имя поля при сериализации/десериализации
	Age         uint8    `json:"age"`
	Course      uint8    `json:"course"`
	Single      bool     `json:"isSingle"` // имя тега может отличаться от имени поля структуры
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
