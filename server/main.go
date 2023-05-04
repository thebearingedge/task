package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thebearingedge/task/application"
)

type App interface {
	FetchRandomNameJoke() (*application.ApplicationResult, error)
}

type Logger interface {
	Err(error)
	Info(args ...any)
}

func NewServer(app App, log Logger) *gin.Engine {
	s := gin.New()
	v1 := s.Group("/v1")
	v1.GET("/random-joke", HandleGetRandomJoke(app, log))
	return s
}
