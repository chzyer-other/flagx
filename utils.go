package reflag

import (
	"errors"
	"fmt"
)

type ErrorFmt struct {
	err  error
	args []interface{}
}

func Error(msg string) *ErrorFmt {
	return &ErrorFmt{err: errors.New(msg)}
}

func (e *ErrorFmt) Format(args ...interface{}) error {
	e.args = args
	return e
}

func (e *ErrorFmt) Error() string {
	return fmt.Sprintf(e.err.Error(), e.args...)
}
