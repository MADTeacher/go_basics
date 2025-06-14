package main

import (
	"encoding/json"
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
	actor := Actor{
		Name:        "Tom Hanks",
		Age:         65,
		FilmsAmount: 50,
		AboutActor:  "Tom Hanks is an actor...",
	}

	actorJson, err := json.MarshalIndent(actor, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	myFile, err := os.Create("actor.json")
	if err != nil {
		log.Fatal(err)
	}

	myFile.Write(actorJson)
	defer myFile.Close()
}
