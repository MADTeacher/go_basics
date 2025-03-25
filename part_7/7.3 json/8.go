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
	group := map[string]any{}

	data, err := os.ReadFile("group.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &group)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", group)
}
