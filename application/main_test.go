package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thebearingedge/task/model"
)

type StubNameFetcher struct {
	stub func() (*model.NameResponse, error)
}

func (n StubNameFetcher) GetRandomName() (*model.NameResponse, error) {
	return n.stub()
}

type StubJokeFetcher struct {
	stub func(firstName string, lastName string) (*model.JokeResponse, error)
}

func (j StubJokeFetcher) GetRandomJoke(firstName string, lastName string) (*model.JokeResponse, error) {
	return j.stub(firstName, lastName)
}

func TestApplicationHandlesFailedNameRequests(t *testing.T) {
	want := assert.AnError
	names := StubNameFetcher{
		stub: func() (*model.NameResponse, error) {
			return nil, want
		},
	}
	jokes := StubJokeFetcher{
		stub: func(firstName, lastName string) (*model.JokeResponse, error) {
			return &model.JokeResponse{}, nil
		},
	}
	app := NewApplication(names, jokes)
	_, got := app.FetchRandomNameJoke()
	assert.Equal(t, want, got)
}

func TestApplicationHandlesFailedJokeRequests(t *testing.T) {
	want := assert.AnError
	names := StubNameFetcher{
		stub: func() (*model.NameResponse, error) {
			return &model.NameResponse{}, nil
		},
	}
	jokes := StubJokeFetcher{
		stub: func(firstName, lastName string) (*model.JokeResponse, error) {
			return nil, want
		},
	}
	app := NewApplication(names, jokes)
	_, got := app.FetchRandomNameJoke()
	assert.Equal(t, want, got)
}

func TestApplicationFetchesCustomJoke(t *testing.T) {
	want := "Rob Pike worked on Go."
	names := StubNameFetcher{
		stub: func() (*model.NameResponse, error) {
			return &model.NameResponse{FirstName: "Rob", LastName: "Pike"}, nil
		},
	}
	jokes := StubJokeFetcher{
		stub: func(firstName, lastName string) (*model.JokeResponse, error) {
			return &model.JokeResponse{
				Type: "success",
				Value: struct {
					Categories []string "json:\"categories\""
					ID         int      "json:\"id\""
					Joke       string   "json:\"joke\""
				}{
					Categories: []string{"nerdy"},
					ID:         42,
					Joke:       firstName + " " + lastName + " worked on Go.",
				},
			}, nil
		},
	}
	app := NewApplication(names, jokes)
	got, err := app.FetchRandomNameJoke()
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
