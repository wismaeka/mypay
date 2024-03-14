// Package validate provides the service for validation.
package validate

import (
	"context"

	"mypayment/business/data/external"
)

type store interface {
	Validate(ctx context.Context, number, name string) (external.Account, error)
}

type Handler struct {
	s store
}

func NewHandler(s store) Handler {
	return Handler{
		s: s,
	}
}

type validateResponse struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	IsValid       bool   `json:"is_valid"`
}
