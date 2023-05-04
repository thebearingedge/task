package log

import (
	"fmt"
	"os"
)

type Logger struct{}

func NewLogger() Logger {
	return Logger{}
}

func (l Logger) Info(args ...any) {
	fmt.Fprintln(os.Stdout, args...)
}

func (l Logger) Err(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}
