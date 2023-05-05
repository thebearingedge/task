package gateway

import (
	"errors"
)

func Retry[T any](op func() (T, error), onErr func(tries int, err error) error) (T, error) {
	var e error
	for tries := 1; ; tries++ {
		if t, err := op(); err != nil {
			e = errors.Join(err, e)
			if err := onErr(tries, err); err != nil {
				e = errors.Join(err, e)
				return t, e
			}
		} else {
			return t, nil
		}
	}
}
