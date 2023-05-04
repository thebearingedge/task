package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thebearingedge/task/application"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type StubApp struct {
	stub func() (*application.ApplicationResult, error)
}

func (a StubApp) FetchRandomNameJoke() (*application.ApplicationResult, error) {
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
		stub: func() (*application.ApplicationResult, error) {
			return nil, want
		},
	}
	l := StubLogger{}
	s := NewServer(a, &l)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/random-joke", nil)
	assert.Nil(t, err)
	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	got := l.errors[0]
	assert.NotNil(t, got)
	assert.Equal(t, want, got)
}

func TestServerResponseWithApplicationResultSuccess(t *testing.T) {
	want := application.ApplicationResult{}
	a := StubApp{
		stub: func() (*application.ApplicationResult, error) {
			return &want, nil
		},
	}
	l := StubLogger{}
	s := NewServer(a, &l)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/random-joke", nil)
	assert.Nil(t, err)
	s.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var got application.ApplicationResult
	json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, want, got)
}
