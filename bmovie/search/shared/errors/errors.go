package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrUnhandledRespStatus = func(status, message interface{}) error { return errors.New(fmt.Sprintf("Unhandled response status: %v, message: %v", status, message)) }
	ErrUnhandledHttpStatus = func(status interface{}) error { return errors.New(fmt.Sprintf("Unhandled http status code: %v", status)) }
)
