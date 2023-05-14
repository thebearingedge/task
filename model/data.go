package model

type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value struct {
		Categories []string `json:"categories"`
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
	} `json:"value"`
}
