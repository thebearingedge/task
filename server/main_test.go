package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thebearingedge/task/application"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type StubApp struct {
	stub func() (application.ApplicationResult, error)
}

func (a StubApp) FetchRandomNameJoke() (application.ApplicationResult, error) {
	return a.stub()
}

func TestXxx(t *testing.T) {
	a := StubApp{}
	assert.NotNil(t, NewServer(a))
}
