package application

import (
	"github.com/thebearingedge/task/model"
)

type NameFetcher interface {
	GetRandomName() (*model.NameResponse, error)
}

type JokeFetcher interface {
	GetRandomJoke(firstName string, lastName string) (*model.JokeResponse, error)
}

type ApplicationResult struct {
	*model.JokeResponse
}

type Application struct {
	names NameFetcher
	jokes JokeFetcher
}

func NewApplication(names NameFetcher, jokes JokeFetcher) Application {
	return Application{names, jokes}
}

func (a Application) FetchRandomNameJoke() (*ApplicationResult, error) {
	name, nameErr := a.names.GetRandomName()
	if nameErr != nil {
		return nil, nameErr
	}
	joke, jokeErr := a.jokes.GetRandomJoke(name.FirstName, name.LastName)
	if jokeErr != nil {
		return nil, jokeErr
	}
	return &ApplicationResult{joke}, nil
}
