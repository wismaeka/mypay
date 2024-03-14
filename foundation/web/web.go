// Package web contains the web framework extension
package web

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Handler is a handler func that:
// 1. receives context so that we can treat the context as a separated entity
// 2. returns error to facilitate centralizing error handling
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers.
type App struct {
	shutdown chan os.Signal
	mux      *mux.Router
}

// ServeHTTP implements the http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(mux *mux.Router) *App {
	return &App{
		mux: mux,
	}
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux
func (a *App) Handle(method, path string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {
		// Pull the context from the request and
		// use it as a separate parameter.
		ctx := r.Context()

		// Capture the parent request from the context.
		traceID := "123456789"
		v := Values{
			TraceID: traceID,
		}
		ctx = WithValues(ctx, &v)
		if err := handler(ctx, w, r); err != nil {
			return
		}
	}
	a.mux.HandleFunc(path, h).Methods(method)
}
