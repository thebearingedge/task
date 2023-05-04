package gateway

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/thebearingedge/task/model"
)

var BaseNameURL string = "https://names.mcquay.me/api/v0/"

type nameHttpClient interface {
	Get(url string) (*http.Response, error)
}
type NameGateway struct {
	client nameHttpClient
	uri    string
}

func NewNameGateway(client nameHttpClient, uri string) NameGateway {
	if _, err := url.ParseRequestURI(uri); err != nil {
		panic(err)
	}
	return NameGateway{client, uri}
}

func (g NameGateway) GetRandomName() (*model.NameResponse, error) {
	var name model.NameResponse
	res, err := g.client.Get(g.uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(data, &name); err != nil {
		return nil, err
	}
	return &name, nil
}
