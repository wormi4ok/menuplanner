package internal

import "errors"

type ErrorFlag int

const (
	ErrorNotFound = iota + 1
	ErrorUnauthorized
)

// NewError wraps err with an error that will return true from ErrorIs(err, flag).
func NewError(err error, flag ErrorFlag) error {
	if err == nil {
		return nil
	}

	return flagErr{error: err, flag: flag}
}

func ErrorIs(err error, flag ErrorFlag) bool {
	for {
		if f, ok := err.(flagErr); ok && f.flag == flag {
			return true
		}
		if err = errors.Unwrap(err); err == nil {
			return false
		}
	}
}

type flagErr struct {
	flag ErrorFlag
	error
}

func (f flagErr) Unwrap() error {
	return f.error
}
