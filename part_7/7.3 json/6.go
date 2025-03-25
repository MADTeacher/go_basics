package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Name        string   `json:"name"`
	Age         uint8    `json:"age"`
	Course      uint8    `json:"course"`
	Single      bool     `json:"isSingle"`
	Description []string `json:"description"`
}

type Group struct {
	GroupName      string    `json:"groupName"`
	Course         uint8     `json:"course"`
	AmountStudents uint8     `json:"amountStudents"`
	Students       []Student `json:"students"`
}

func main() {
	students := []Student{
		{
			Name:   "Alex",
			Age:    20,
			Course: 2,
			Single: true,
			Description: []string{
				"Мечтатель",
				"Ленив",
				"Постоянно жалуется на жизнь",
			},
		},
		{
			Name:   "Max",
			Age:    16,
			Course: 2,
			Single: true,
			Description: []string{
				"Реалист",
				"Мега-мозг",
			},
		},
		{
			Name:   "Smith",
			Age:    20,
			Course: 2,
			Single: false,
			Description: []string{
				"Паникер",
				"Двоечник",
				"Девиз: жизнь - боль...",
			},
		},
	}

	group := Group{
		GroupName:      "4646M1",
		Course:         2,
		AmountStudents: uint8(len(students)),
		Students:       students,
	}

	data, err := json.MarshalIndent(&group, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	myFile, err := os.Create("group.json")
	if err != nil {
		log.Fatal(err)
	}
	myFile.Write(data)
	defer myFile.Close()

}
