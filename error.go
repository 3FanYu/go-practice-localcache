package localcache

import (
	"errors"
	"fmt"
)

const (
	KeyNotFound = "KeyNotFound"
)

type Error struct {
	err  error
	k    string
	code string
}

func (e *Error) Error() string {
	return e.err.Error()
}

func NewKeyNotFoundError(k string) error {
	err := errors.New(fmt.Sprintf("key %s is not found", k))
	return &Error{
		k:    k,
		err:  err,
		code: KeyNotFound,
	}
}
