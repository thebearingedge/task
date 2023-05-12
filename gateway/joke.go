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

// TODO - returning a pointer may not be desirable
// however, representing absence without nil relies on meaningful errors
// and i'm not sure returning a "zero-value" struct is actually a good idea
func (g jokeGateway) GetRandomJoke(firstName string, lastName string) (*model.JokeResponse, error) {
	query := url.Values{}
	query.Set("firstName", firstName)
	query.Set("lastName", lastName)
	url := fmt.Sprintf("%v?%v", g.uri, query.Encode())
	res, err := g.client.Get(url)
	// TODO - the http Client does not return an error
	// for 4xx and 5xx responses
	// there should be test cases for this failure mode
	if err != nil {
		return nil, fmt.Errorf("send request to %v for joke: %w", g.uri, err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	var joke model.JokeResponse
	if err := json.Unmarshal(data, &joke); err != nil {
		return nil, fmt.Errorf("unmarshal joke response data: - %w", err)
	}
	return &joke, nil
}
