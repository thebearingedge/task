package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thebearingedge/task/application"
)

type App interface {
	FetchRandomNameJoke() (application.ApplicationResult, error)
}

func NewServer(app App) *gin.Engine {
	s := gin.New()
	return s
}
