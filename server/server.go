package server

import (
	"github.com/gin-gonic/gin"
)

type App interface {
	FetchRandomNameJoke() (string, error)
}

type Logger interface {
	Err(error)
	Info(args ...any)
}

// TODO - it would be nice to track failures
// and incorporate some kind of circuit breaker
// but i'll need to learn about how that affects throughput
func NewServer(app App, log Logger) *gin.Engine {
	s := gin.New()
	v1 := s.Group("/v1")
	v1.GET("/random-joke", HandleGetRandomJoke(app, log))
	return s
}
