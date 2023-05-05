package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubNamesHttpClient struct {
	stub func(url string) (*http.Response, error)
}

func (j StubNamesHttpClient) Get(url string) (*http.Response, error) {
	return j.stub(url)
}

func TestCannotMakeNamesGatewayWithInvalidURI(t *testing.T) {
	assert.Panics(t, func() {
		NewNameGateway(http.DefaultClient, "nope")
	})
}

func TestGetNameHttpError(t *testing.T) {
	want := assert.AnError
	s := StubNamesHttpClient{
		stub: func(url string) (*http.Response, error) {
			return nil, want
		},
	}
	g := NewNameGateway(s, "https://names.mcquay.me/api/v0/")
	name, got := g.GetRandomName()
	assert.Nil(t, name)
	assert.ErrorIs(t, got, want)
}

func TestGetNameClientError(t *testing.T) {
	var want gatewayError
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer s.Close()
	g := NewNameGateway(s.Client(), s.URL)
	name, err := g.GetRandomName()
	assert.Nil(t, name)
	assert.ErrorAs(t, err, &want)
	assert.ErrorContains(t, err, "client error")
}

func TestGetNameServerError(t *testing.T) {
	var want gatewayError
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer s.Close()
	g := NewNameGateway(s.Client(), s.URL)
	name, err := g.GetRandomName()
	assert.Nil(t, name)
	assert.ErrorAs(t, err, &want)
	assert.ErrorContains(t, err, "server error")
}

func TestGetNameUnmarshalError(t *testing.T) {
	var want *json.SyntaxError
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()
	g := NewNameGateway(s.Client(), s.URL)
	name, got := g.GetRandomName()
	assert.Nil(t, name)
	assert.ErrorAs(t, got, &want)
}

func TestGetNameSuccess(t *testing.T) {
	nameFixtureBytes, err := ioutil.ReadFile("../fixture/name.json")
	assert.Nil(t, err)
	want := strings.TrimSpace(string(nameFixtureBytes))
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, want)
	}))
	defer s.Close()
	g := NewNameGateway(s.Client(), s.URL)
	name, err := g.GetRandomName()
	assert.Nil(t, err)
	nameResponseBytes, err := json.Marshal(name)
	assert.Nil(t, err)
	got := strings.TrimSpace(string(nameResponseBytes))
	assert.Equal(t, want, got)
}
