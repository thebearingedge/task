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

type StubJokesHttpClient struct {
	stub func(url string) (*http.Response, error)
}

func (j StubJokesHttpClient) Get(url string) (*http.Response, error) {
	return j.stub(url)
}

func TestCannotMakeJokesGatewayWithInvalidURI(t *testing.T) {
	assert.Panics(t, func() {
		NewJokeGateway(http.DefaultClient, "nope")
	})
}

func TestGetJokeResponseErrorHTTP(t *testing.T) {
	want := assert.AnError
	s := StubJokesHttpClient{
		stub: func(url string) (*http.Response, error) {
			return nil, want
		},
	}
	g := NewJokeGateway(s, BaseJokeURL)
	joke, got := g.GetRandomJoke("Rob", "Pike")
	assert.Nil(t, joke)
	assert.ErrorContains(t, got, want.Error())
}

func TestGetJokeResponseErrorUnmarshal(t *testing.T) {
	want := json.SyntaxError{}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()
	g := NewJokeGateway(s.Client(), s.URL)
	joke, got := g.GetRandomJoke("Rob", "Pike")
	assert.Nil(t, joke)
	assert.ErrorContains(t, got, want.Error())
}

func TestGetJokeResponseSuccess(t *testing.T) {
	jokeFixtureBytes, err := ioutil.ReadFile("../fixture/joke.json")
	assert.Nil(t, err)
	want := strings.TrimSpace(string(jokeFixtureBytes))
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, want)
	}))
	defer s.Close()
	g := NewJokeGateway(s.Client(), s.URL)
	joke, err := g.GetRandomJoke("Rob", "Pike")
	assert.Nil(t, err)
	jokeResponseBytes, err := json.Marshal(joke)
	assert.Nil(t, err)
	got := strings.TrimSpace(string(jokeResponseBytes))
	assert.Equal(t, want, got)
}
