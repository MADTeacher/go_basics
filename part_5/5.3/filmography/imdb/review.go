package imdb

type Review struct {
	Name   string `json:"name"`
	Text   string `json:"text"`
	Rating int    `json:"rating"`
}
