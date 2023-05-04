package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/thebearingedge/task/model"
)

var BaseJokeURL string = "http://joke.loc8u.com:8888/joke"

type jokeHttpClient interface {
	Get(url string) (*http.Response, error)
}

type JokeGateway struct {
	client jokeHttpClient
	uri    string
}

func NewJokeGateway(client jokeHttpClient, uri string) JokeGateway {
	if _, err := url.ParseRequestURI(uri); err != nil {
		panic(err)
	}
	return JokeGateway{client, uri}
}

func (g JokeGateway) GetRandomJoke(firstName string, lastName string) (*model.JokeResponse, error) {
	var joke model.JokeResponse
	url := fmt.Sprintf("%s?limitTo=nerdy&firstName=%s&lastName=%s", g.uri, firstName, lastName)
	res, err := g.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(data, &joke); err != nil {
		return nil, err
	}
	return &joke, nil
}
