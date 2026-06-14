package main

import (
	"encoding/json"
	"fmt"
	"log"
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
        Age:         69,
        FilmsAmount: 70,
		AboutActor:  "Tom Hanks is an actor...",
	}

	actorJson, err := json.Marshal(actor)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(actorJson))
}
