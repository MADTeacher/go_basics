package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Actor struct {
	// имя поля при сериализации/десериализации
	Name string `json:"name"`
	Age  int    `json:"age"`
	// имя тега может отличаться от имени поля структуры
	FilmsAmount int     `json:"films_amount,omitempty"`
	AboutActor  *string `json:",omitempty"`
}

func main() {
	actorJson := map[string]any{}

	data, err := os.ReadFile("actor.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &actorJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", actorJson)
}
