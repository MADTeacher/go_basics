package imdb

type Actor struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	FilmsAmount int    `json:"filmsAmount"`
	AboutActor  string `json:"aboutActor"`
}
