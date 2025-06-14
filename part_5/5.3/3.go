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
	AboutActor  string
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

	fmt.Printf("%+v", *actor)
}
