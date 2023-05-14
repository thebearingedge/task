package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type StubApp struct {
	stub func() (string, error)
}

func (a StubApp) FetchRandomNameJoke() (string, error) {
	return a.stub()
}

type StubLogger struct {
	stubInfo func(args ...any)
	errors   []error
}

func (l *StubLogger) Info(args ...any) {}
func (l *StubLogger) Err(err error) {
	l.errors = append(l.errors, err)
}

func TestServerResponseWithApplicationResultFail(t *testing.T) {
	want := assert.AnError
	a := StubApp{
		stub: func() (string, error) {
			return "", want
		},
	}
	l := StubLogger{}
	s := NewServer(a, &l)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/random-joke", nil)
	assert.Nil(t, err)
	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Len(t, l.errors, 1)
	got := l.errors[0]
	assert.NotNil(t, got)
	assert.Equal(t, want, got)
}

func TestServerResponseWithApplicationResultSuccess(t *testing.T) {
	want := `
		There are two hard things in distributed systems:
		2. exactly-once delivery
		1. in-order delivery
		2. exactly-once delivery
	`
	a := StubApp{
		stub: func() (string, error) {
			return want, nil
		},
	}
	l := StubLogger{}
	s := NewServer(a, &l)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/random-joke", nil)
	assert.Nil(t, err)
	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	got := w.Body.String()
	assert.Equal(t, want, got)
}
