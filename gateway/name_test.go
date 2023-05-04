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

func TestGetNameResponseErrorHTTP(t *testing.T) {
	want := assert.AnError
	s := StubNamesHttpClient{
		stub: func(url string) (*http.Response, error) {
			return nil, want
		},
	}
	g := NewNameGateway(s, "https://names.mcquay.me/api/v0/")
	name, got := g.GetRandomName()
	assert.Nil(t, name)
	assert.ErrorContains(t, got, want.Error())
}

func TestGetNameResponseErrorUnmarshal(t *testing.T) {
	want := json.SyntaxError{}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()
	g := NewNameGateway(s.Client(), s.URL)
	name, got := g.GetRandomName()
	assert.Nil(t, name)
	assert.ErrorContains(t, got, want.Error())
}

func TestGetNameResponseSuccess(t *testing.T) {
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
