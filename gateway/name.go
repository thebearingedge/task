package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/thebearingedge/task/model"
)

var BaseNameURL string = "https://names.mcquay.me/api/v0/"

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
	var name model.NameResponse
	res, err := g.client.Get(g.uri)
	if err != nil {
		return nil, fmt.Errorf("send request to %v for name: %w", g.uri, err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(data, &name); err != nil {
		return nil, fmt.Errorf("unmarshal name response data: - %w", err)
	}
	return &name, nil
}
