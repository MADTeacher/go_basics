package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Actor struct {
	// имя поля при сериализации/десериализации
	Name string `json:"name"`
	Age  int    `json:"age"`
	// имя тега может отличаться от имени поля структуры
	FilmsAmount int     `json:"films_amount"`
	AboutActor  *string `json:"about"`
}

func main() {
	aboutStr := "Tom Hanks is an actor..."
	actor := Actor{
		Name:        "Tom Hanks",
		Age:         65,
		FilmsAmount: 50,
		AboutActor:  &aboutStr,
	}

	actorJson, err := json.MarshalIndent(actor, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(actorJson))
}
