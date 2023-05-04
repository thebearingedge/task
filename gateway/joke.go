package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/thebearingedge/task/model"
)

type jokeHttpClient interface {
	Get(url string) (*http.Response, error)
}

type jokeGateway struct {
	client jokeHttpClient
	uri    string
}

func NewJokeGateway(client jokeHttpClient, uri string) jokeGateway {
	if _, err := url.ParseRequestURI(uri); err != nil {
		panic(err)
	}
	return jokeGateway{client, uri}
}

func (g jokeGateway) GetRandomJoke(firstName string, lastName string) (*model.JokeResponse, error) {
	var joke model.JokeResponse
	url := fmt.Sprintf("%s?limitTo=nerdy&firstName=%s&lastName=%s", g.uri, firstName, lastName)
	res, err := g.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("send request to %v for joke: %w", g.uri, err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(data, &joke); err != nil {
		return nil, fmt.Errorf("unmarshal joke response data: - %w", err)
	}
	return &joke, nil
}
