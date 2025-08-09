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
	FilmsAmount int     `json:"films_amount,omitempty"`
	AboutActor  *string `json:",omitempty"`
}

func ActorToBytes(actor Actor) []byte {
	actorJson, err := json.MarshalIndent(actor, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return actorJson
}

func main() {
	aboutStr := "Tom Hanks is an actor..."
	// поле FilmsAmount проинициализируется значением
	// по умолчанию
	actor := Actor{
		Name:       "Tom Hanks",
		Age:        65,
		AboutActor: &aboutStr,
	}

	actorJson := ActorToBytes(actor)
	// Выводим сериализованные данные
	fmt.Println()
	fmt.Println("First serialized data:")
	fmt.Println(string(actorJson))

	// поле AboutActor и Age проинициализируются значением
	// по умолчанию
	actor = Actor{
		Name:        "Tom Hanks",
		FilmsAmount: 150,
	}

	actorJson = ActorToBytes(actor)
	// Выводим сериализованные данные
	fmt.Println()
	fmt.Println("Second serialized data:")
	fmt.Println(string(actorJson))
}
