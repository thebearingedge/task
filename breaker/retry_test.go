package breaker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetryStrategyAbortsOnError(t *testing.T) {
	want := errors.New("max retries exceeded")
	op := func() error {
		return assert.AnError
	}
	strat := func(retry int, err error) error {
		assert.ErrorIs(t, err, assert.AnError)
		if retry == 3 {
			return want
		}
		return nil
	}
	got := Retry(op, strat)
	assert.ErrorIs(t, got, want)
	assert.ErrorIs(t, got, assert.AnError)
}

func TestRetryStrategyIsSkippedOnSuccess(t *testing.T) {
	op := func() error {
		return nil
	}
	strat := func(retry int, err error) error {
		assert.FailNow(t, "retry strategy should not be called")
		return nil
	}
	err := Retry(op, strat)
	assert.Nil(t, err)
}

func TestRetryOperationUntilSucceeded(t *testing.T) {
	try := 0
	op := func() error {
		if try == 3 {
			return nil
		}
		try += 1
		return assert.AnError
	}
	strat := func(retry int, err error) error {
		assert.LessOrEqual(t, retry, 3)
		return nil
	}
	err := Retry(op, strat)
	assert.Nil(t, err)
}
