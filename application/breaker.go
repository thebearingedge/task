package application

import (
	"sync"
	"time"
)

type breakerState int

const (
	closed breakerState = iota
	open
	halfOpen
)

type BreakerConfig struct {
	OpenThreshold    int
	CloseThreshold   int
	HalfOpenTryOneIn int
	HalfOpenAfter    time.Duration
}

type breaker struct {
	sync.Mutex
	state            breakerState
	failures         int
	openThreshold    int
	closeThreshold   int
	halfOpenAfter    time.Duration
	halfOpenTryOneIn int
	oneIn            int
}

func (b *breaker) IsOpen() bool {
	b.Lock()
	defer b.Unlock()
	if b.state == halfOpen {
		b.oneIn = (b.oneIn + 1) % b.halfOpenTryOneIn
		return b.oneIn != 0
	}
	return b.state == open
}

func (b *breaker) Pass() {
	b.Lock()
	defer b.Unlock()
	if b.state == closed {
		return
	}
	b.failures = min(0, b.failures-1)
	if b.openThreshold-b.failures <= b.closeThreshold {
		b.failures = 0
		b.state = closed
	}
}

func (b *breaker) Fail() {
	b.Lock()
	defer b.Unlock()
	b.failures = min(b.openThreshold, b.failures+1)
	if b.failures != b.openThreshold || b.state == open {
		return
	}
	b.state = open
	go func(halfOpenIn time.Duration) {
		<-time.After(halfOpenIn)
		b.Lock()
		defer b.Unlock()
		b.state = halfOpen
	}(b.halfOpenAfter)
}

func NewBreaker(c BreakerConfig) breaker {
	closeThreshold := c.CloseThreshold
	if c.CloseThreshold == 0 {
		closeThreshold = c.OpenThreshold
	}
	halfOpenTryOneIn := c.HalfOpenTryOneIn
	if halfOpenTryOneIn == 0 {
		halfOpenTryOneIn = 10
	}
	return breaker{
		state:            closed,
		openThreshold:    c.OpenThreshold,
		closeThreshold:   closeThreshold,
		halfOpenAfter:    c.HalfOpenAfter,
		halfOpenTryOneIn: c.HalfOpenTryOneIn,
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
