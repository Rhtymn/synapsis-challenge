package apperror

import (
	"fmt"
	"runtime/debug"
)

type AppError struct {
	Code    int
	Message string
	err     error
	stack   []byte
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		err:     err,
		stack:   debug.Stack(),
	}
}

func (e *AppError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("(%d) %s: %s", e.Code, e.Message, e.err)
	}
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) GetStackTrace() []byte {
	return e.stack
}

func (e *AppError) ContainsStackTrace() bool {
	return len(e.stack) > 0
}

func IsErrorCode(err error, code int) bool {
	aerr, ok := err.(*AppError)
	return ok && aerr.Code == code
}

func Wrap(err error) error {
	if _, ok := err.(*AppError); ok {
		return err
	}
	return NewInternal(err)
}

func NewInternal(err error) error {
	return NewAppError(CodeInternal, "internal error", err)
}

func NewInternalFmt(format string, args ...interface{}) error {
	return NewInternal(fmt.Errorf(format, args...))
}

func NewTypeAssertionFailed(want interface{}, got interface{}) error {
	return NewInternalFmt("type assert: want %T, got %T", want, got)
}

func NewBadRequest(err error, message string) error {
	return NewAppError(
		CodeBadRequest,
		message,
		err,
	)
}

func NewNotFound(err error, message string) error {
	return NewAppError(
		CodeNotFound,
		message,
		err,
	)
}

func NewAlreadyExists(err error, message string) error {
	return NewAppError(
		CodeAlreadyExists,
		message,
		err,
	)
}

func NewUnauthorized(err error, message string) error {
	return NewAppError(
		CodeUnauthorized,
		message,
		err,
	)
}

func NewForbidden(err error, message string) error {
	return NewAppError(
		CodeForbidden,
		message,
		err,
	)
}