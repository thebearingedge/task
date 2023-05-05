package gateway

import (
	"fmt"
)

type gatewayErrorKind int

const (
	httpError gatewayErrorKind = iota
	clientError
	serverError
	deserializationError
)

type gatewayError struct {
	kind gatewayErrorKind
	err  error
}

func (e gatewayError) Error() string {
	return e.err.Error()
}

func (e gatewayError) Unwrap() error {
	return e.err
}

func httpErrorf(format string, args ...any) gatewayError {
	return gatewayError{
		httpError,
		fmt.Errorf(format, args...),
	}
}

func clientErrorf(format string, args ...any) gatewayError {
	return gatewayError{
		clientError,
		fmt.Errorf(format, args...),
	}
}

func serverErrorf(format string, args ...any) gatewayError {
	return gatewayError{
		serverError,
		fmt.Errorf(format, args...),
	}
}

func deserializationErrorf(format string, args ...any) gatewayError {
	return gatewayError{
		deserializationError,
		fmt.Errorf(format, args...),
	}
}
