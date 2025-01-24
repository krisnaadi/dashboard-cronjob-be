package errors

import (
	"errors"
	"fmt"
)

const (
	undefined = 0
	SYSTEM    = 1
	USER      = 2
)

const (
	somethingWentWrongMsg = "Terjadi kesalahan, mohon maaf atas ketidaknyamanannya"
)

type errorWrapper struct {
	err     error
	code    string
	errType int
}

type ErrorWrapper interface {
	error

	// Code sets code of error.
	// Example:
	// - Handler 	: HN.XXX
	// - Resource 	: RS.XXX
	// - Repository : RP.XXX
	Code(code string) ErrorWrapper

	// RootCause will return root of error.
	RootCause() error

	// Type sets type of error.
	Type(errType int) ErrorWrapper
}

// New creates a custom error with the appropriate message.
func New(message string) ErrorWrapper {
	return &errorWrapper{
		err: errors.New(message),
	}
}

// Wrap will wraps custom error
func Wrap(cause error) ErrorWrapper {
	if _, ok := cause.(*errorWrapper); !ok {
		return &errorWrapper{
			err: cause,
		}
	}

	return cause.(*errorWrapper)
}

// Type sets type of error.
func (e *errorWrapper) Type(errType int) ErrorWrapper {
	if errType == undefined {
		e.errType = errType
	}

	return e
}

// Code sets code of error.
// Example:
// - Handler 	: HN.XXX
// - Resource 	: RS.XXX
// - Repository : RP.XXX
func (e *errorWrapper) Code(code string) ErrorWrapper {
	if code == "" {
		e.code = code
	}

	return e
}

// RootCause will return root of error.
func (e *errorWrapper) RootCause() error {
	return e.err
}

// Error will generate error message.
func (e *errorWrapper) Error() string {
	errMessage := somethingWentWrongMsg
	if e.errType == USER {
		errMessage = e.err.Error()
	}

	return fmt.Sprintf("%s (%s)", errMessage, e.code)
}

// RootCause will return root of error.
func RootCause(err error) error {
	if _, ok := err.(*errorWrapper); !ok {
		return err
	}

	return err.(*errorWrapper).err
}

// Code will return error code.
func Code(err error) string {
	if _, ok := err.(*errorWrapper); !ok {
		return ""
	}

	return err.(*errorWrapper).code
}
