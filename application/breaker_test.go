package application

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBreakerOpensAtOpenThreshold(t *testing.T) {
	c := BreakerConfig{OpenThreshold: 3, HalfOpenAfter: time.Millisecond}
	b := NewBreaker(c)
	assert.False(t, b.IsOpen())
	b.Fail()
	assert.False(t, b.IsOpen())
	b.Fail()
	assert.False(t, b.IsOpen())
	b.Fail()
	assert.True(t, b.IsOpen())
}

func TestBreakerClosesAtCloseThreshold(t *testing.T) {
	c := BreakerConfig{OpenThreshold: 3, HalfOpenAfter: time.Millisecond}
	b := NewBreaker(c)
	b.Fail()
	b.Fail()
	b.Fail()
	assert.True(t, b.IsOpen())
	b.Pass()
	b.Pass()
	b.Pass()
	assert.False(t, b.IsOpen())
}

func TestBreakerIsRegularlyClosedWhileHalfOpen(t *testing.T) {
	c := BreakerConfig{OpenThreshold: 3, HalfOpenAfter: time.Millisecond, HalfOpenTryOneIn: 3}
	b := NewBreaker(c)
	b.Fail()
	b.Fail()
	b.Fail()
	assert.True(t, b.IsOpen())
	<-time.After(2 * time.Millisecond)
	assert.True(t, b.IsOpen())
	assert.True(t, b.IsOpen())
	assert.False(t, b.IsOpen())
	assert.True(t, b.IsOpen())
	assert.True(t, b.IsOpen())
	assert.False(t, b.IsOpen())
}
