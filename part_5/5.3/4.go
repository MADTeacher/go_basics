package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Actor struct {
	Name        string
	Age         int
	FilmsAmount int
	AboutActor  *string
}

func main() {
	actor := &Actor{}

	data, err := os.ReadFile("actor.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, actor)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим данные, прочитанные из файла
	fmt.Print("Deserialized data: ")
	fmt.Printf("%+v\n", *actor)

	actorJson, err := json.MarshalIndent(actor, "", "  ") // сериализация
	if err != nil {
		log.Fatal(err)
	}

	// Выводим сериализованные данные
	fmt.Println()
	fmt.Println("Serialized data:")
	fmt.Println(string(actorJson))
}
