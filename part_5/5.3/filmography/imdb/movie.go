package imdb

type Movie struct {
	Name           string   `json:"name"`
	Budget         int      `json:"budget"`
	Actors         []Actor  `json:"actors"`
	CriticsRating  float64  `json:"criticsRating"`
	AudienceRating float64  `json:"audienceRating"`
	Year           int      `json:"year"`
	Country        string   `json:"country"`
	Genre          Genre    `json:"genre"`
	Reviews        []Review `json:"reviews"`
}
