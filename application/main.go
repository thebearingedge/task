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

type Application struct {
	names NameFetcher
	jokes JokeFetcher
}

func NewApplication(names NameFetcher, jokes JokeFetcher) Application {
	return Application{names, jokes}
}

func (a Application) FetchRandomNameJoke() (string, error) {
	name, nameErr := a.names.GetRandomName()
	if nameErr != nil {
		return "", nameErr
	}
	joke, jokeErr := a.jokes.GetRandomJoke(name.FirstName, name.LastName)
	if jokeErr != nil {
		return "", jokeErr
	}
	return joke.Value.Joke, nil
}
