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

var nameFixtureBytes, _ = ioutil.ReadFile("../fixture/name.json")

func TestSuccessfulGetNameResponse(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, strings.TrimSpace(string(nameFixtureBytes)))
	}))
	defer s.Close()
	g := NewNameGateway(s.Client(), s.URL)
	name, err := g.GetRandomName()
	assert.Nil(t, err)
	nameResponseBytes, err := json.Marshal(name)
	assert.Nil(t, err)
	assert.Equal(
		t,
		strings.TrimSpace(string(nameFixtureBytes)),
		strings.TrimSpace(string(nameResponseBytes)),
	)
}
