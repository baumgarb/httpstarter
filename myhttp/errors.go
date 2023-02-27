package myhttp

import "fmt"

type ClientErr struct {
	inner error
}

func (err ClientErr) Error() string {
	return fmt.Sprintf("client error: %v", err.inner)
}

func NewClientError(err error) error {
	return ClientErr{
		inner: err,
	}
}

func NewClientErrorf(format string, a ...any) error {
	return ClientErr{
		inner: fmt.Errorf(format, a...),
	}
}
