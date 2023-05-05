package breaker

import (
	"errors"
)

type operation func() error

type strategy func(try int, err error) error

func Retry(op operation, onErr strategy) error {
	var e error
	done := make(chan bool, 1)
	fail := make(chan error, 1)
	go doRetry(1, op, onErr, fail, done)
	for err := range fail {
		e = errors.Join(e, err)
	}
	if <-done {
		return nil
	}
	return e
}

func doRetry(retry int, op operation, onRetry strategy, fail chan<- error, done chan<- bool) {
	err := op()
	if err == nil {
		done <- true
		close(fail)
		close(done)
		return
	}
	fail <- err
	if dead := onRetry(retry, err); dead != nil {
		fail <- dead
		done <- false
		close(fail)
		close(done)
		return
	}
	doRetry(retry+1, op, onRetry, fail, done)
}
