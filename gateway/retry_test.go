package gateway

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetryStrategyAbortsOnError(t *testing.T) {
	want := errors.New("max retries exceeded")
	op := func() (any, error) {
		return nil, assert.AnError
	}
	strat := func(retry int, err error) error {
		assert.ErrorIs(t, err, assert.AnError)
		if retry == 3 {
			return want
		}
		return nil
	}
	_, got := Retry(op, strat)
	assert.ErrorIs(t, got, want)
	assert.ErrorIs(t, got, assert.AnError)
}

func TestRetryStrategyIsSkippedOnSuccess(t *testing.T) {
	want := 42
	op := func() (int, error) {
		return want, nil
	}
	strat := func(retry int, err error) error {
		assert.FailNow(t, "retry strategy should not be called")
		return nil
	}
	got, err := Retry(op, strat)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestRetryOperationUntilSucceeded(t *testing.T) {
	try := 0
	var want int
	op := func() (int, error) {
		if try == 3 {
			return want, nil
		}
		try += 1
		return want, assert.AnError
	}
	strat := func(retry int, err error) error {
		assert.LessOrEqual(t, retry, 3)
		return nil
	}
	got, err := Retry(op, strat)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
