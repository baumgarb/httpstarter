package httpstarter

import (
	"errors"
	"fmt"
)

var ErrEntityNotFound = errors.New("entity not found")

type InputValidationError struct {
	msg string
}

func (e InputValidationError) Error() string {
	return e.msg
}

func NewInputValidationError(format string, a ...any) InputValidationError {
	return InputValidationError{
		msg: fmt.Sprintf(format, a...),
	}
}
