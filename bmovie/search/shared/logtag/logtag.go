package logtag

const (
	PerformSearchMovie = "search movie"

	RequestTmpl = "receive request to %s."
	ErrorTmpl   = "an error occurred when try to %s. ERROR:%v"

	ErrBuildHttpRequest    = "build http request"
	ErrDoHTTPCall          = "do http call"
	ErrReadResponseBody    = "read response body"
	ErrUnhandledHTTPStatus = "handle http status"
	ErrRepositoryStore     = "store record(s)"
)
