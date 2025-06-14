package main

import (
	"encoding/json"
	"filmography/imdb"
	"fmt"
	"log"
	"os"
)

func main() {
	movie := &imdb.Movie{}
	data, err := os.ReadFile("movie.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", movie)

	// записать в файл output.json
	data, err = json.MarshalIndent(movie, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("output.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
