// Package external provides the service for external data.
package external

import (
	"net/http"
	"net/url"
)

// BankCall contains everything it needs to perform call to Bank.
type BankCall struct {
	client  *http.Client
	baseURL *url.URL
}

// New returns a new ValidateLookup.
func New(client *http.Client, baseURL *url.URL) BankCall {
	return BankCall{
		client:  client,
		baseURL: baseURL,
	}
}
