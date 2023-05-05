package gateway

import (
	"encoding/json"
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

// TODO - returning a pointer may not be desirable
// however, representing absence without nil relies on meaningful errors
// and i'm not sure returning a "zero-value" struct is actually a good idea
func (g nameGateway) GetRandomName() (*model.NameResponse, error) {
	res, err := g.client.Get(g.uri)
	// TODO - the http Client does not return an error
	// for 4xx and 5xx responses
	// there should be test cases for this failure mode
	if err != nil {
		return nil, httpErrorf("send request to %v for name: %w", g.uri, err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	var name model.NameResponse
	if err := json.Unmarshal(data, &name); err != nil {
		return nil, deserializationErrorf("unmarshal name response data: - %w", err)
	}
	return &name, nil
}
