package web

import (
	"context"
	"errors"
	"net/http"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// key is how request values are stored/retrieved.
const key ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID        string
	StatusCode     int
	ResponseBody   []byte
	ResponseHeader map[string][]string
}

// WithValues returns a new context with values.
func WithValues(ctx context.Context, val *Values) context.Context {
	return context.WithValue(ctx, key, val)
}

// GetValues returns the values from the context.
func GetValues(ctx context.Context) (*Values, error) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return nil, errors.New("web value missing from context")
	}
	return v, nil
}

// SetStatusCode sets the status code back into the context.
func SetStatusCode(ctx context.Context, statusCode int) error {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return errors.New("web value missing from context")
	}
	v.StatusCode = statusCode
	return nil
}

// SetResponseBody sets the response body into context, for logging purpose
func SetResponseBody(ctx context.Context, resp []byte) error {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return errors.New("web value missing from context")
	}
	v.ResponseBody = resp
	return nil
}

// SetResponseHeader sets the response headers into context, for logging purpose
func SetResponseHeader(ctx context.Context, headers http.Header) error {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return errors.New("web value missing from context")
	}
	v.ResponseHeader = headers
	return nil
}
