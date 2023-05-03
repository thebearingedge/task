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

var jokeFixtureBytes, _ = ioutil.ReadFile("../fixture/joke.json")

func TestSuccessfulGetJokeResponse(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, strings.TrimSpace(string(jokeFixtureBytes)))
	}))
	defer s.Close()
	g := NewJokeGateway(s.Client(), s.URL)
	name, err := g.GetRandomJoke("Rob", "Pike")
	assert.Nil(t, err)
	nameResponseBytes, err := json.Marshal(name)
	assert.Nil(t, err)
	assert.Equal(
		t,
		strings.TrimSpace(string(jokeFixtureBytes)),
		strings.TrimSpace(string(nameResponseBytes)),
	)
}
