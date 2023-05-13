package log

import (
	"fmt"
	"os"
)

// TODO - This logger may not be necessary
// logging in gin: https://gin-gonic.com/docs/examples/custom-log-format/
// custom middleware: https://gin-gonic.com/docs/examples/custom-middleware/
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
