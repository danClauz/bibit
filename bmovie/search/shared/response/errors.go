package response

import "strings"

type (
	StdError interface {
		StdResp
		AppendError(err error)
		GetErrors() []string
		Error() string
	}

	stdError struct {
		stat   status
		errors []string
	}
)

func (e *stdError) GetCode() string {
	return e.stat.GetCode()
}

func (e *stdError) GetStatus() string {
	return e.stat.GetStatus()
}

func (e *stdError) GetMessage() string {
	return e.stat.GetMessage()
}

func (e *stdError) GetHttpStatus() int {
	return e.stat.GetHttpStatus()
}

func (e *stdError) AppendError(err error) {
	if err != nil {
		e.errors = append(e.errors, err.Error())
	}
}

func (e *stdError) GetErrors() []string {
	return e.errors
}

func (e *stdError) Error() string {
	return strings.Join(e.errors, "\n")
}

func Error(stat status, err error) StdError {
	stdErr := stdError{
		stat: stat,
	}

	stdErr.AppendError(err)

	return &stdErr
}
