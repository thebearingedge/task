package gateway

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/thebearingedge/task/model"
)

type nameHttpClient interface {
	Get(url string) (*http.Response, error)
}

type nameGateway struct {
	client nameHttpClient
	uri    string
}

func NewNameGateway(client nameHttpClient, uri string) nameGateway {
	if _, err := url.ParseRequestURI(uri); err != nil {
		panic(err)
	}
	return nameGateway{client, uri}
}

func (g nameGateway) GetRandomName() (*model.NameResponse, error) {
	res, err := g.client.Get(g.uri)
	if err != nil {
		return nil, httpErrorf("malformed request to %v for name: %w", g.uri, err)
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
	var name model.NameResponse
	if err := json.Unmarshal(data, &name); err != nil {
		return nil, deserializationErrorf("unmarshal name response data: - %w", err)
	}
	return &name, nil
}
