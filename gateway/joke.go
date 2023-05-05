package gateway

import (
	"encoding/json"
	"errors"
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
	query := url.Values{}
	query.Set("firstName", firstName)
	query.Set("lastName", lastName)
	url := fmt.Sprintf("%v?%v", g.uri, query.Encode())
	res, err := g.client.Get(url)
	if err != nil {
		return nil, httpErrorf("malformed request to %v for joke: %w", g.uri, err)
	}
	if res.StatusCode >= 500 {
		return nil, serverErrorf(
			"got %v from %v: %w",
			res.Status,
			g.uri,
			errors.New("server error"),
		)
	}
	if res.StatusCode >= 400 {
		return nil, clientErrorf(
			"got %v from %v: %w",
			res.Status,
			g.uri,
			errors.New("client error"),
		)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	var joke model.JokeResponse
	if err := json.Unmarshal(data, &joke); err != nil {
		return nil, deserializationErrorf("unmarshal joke response data: - %w", err)
	}
	return &joke, nil
}
