package gateway

import (
	"fmt"
)

type GatewayErrorKind int

const (
	httpError GatewayErrorKind = iota
	clientError
	serverError
	deserializationError
)

type GatewayError struct {
	err  error
	kind GatewayErrorKind
}

func (e GatewayError) Error() string {
	return e.err.Error()
}

func (e GatewayError) Unwrap() error {
	return e.err
}

func httpErrorf(format string, args ...any) GatewayError {
	return GatewayError{
		fmt.Errorf(format, args...),
		httpError,
	}
}

func clientErrorf(format string, args ...any) GatewayError {
	return GatewayError{
		fmt.Errorf(format, args...),
		clientError,
	}
}

func serverErrorf(format string, args ...any) GatewayError {
	return GatewayError{
		fmt.Errorf(format, args...),
		serverError,
	}
}

func deserializationErrorf(format string, args ...any) GatewayError {
	return GatewayError{
		fmt.Errorf(format, args...),
		deserializationError,
	}
}
