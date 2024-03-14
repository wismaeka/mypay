package validate

import (
	"context"
	"errors"
	"net/http"

	"mypayment/business/data/external"
	"mypayment/foundation/web"
)

func (h Handler) Account(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	bankNumber := r.URL.Query().Get("bank_number")

	result, err := h.s.Validate(ctx, bankNumber, name)
	if errors.Is(err, external.ErrNotFound) {
		return web.Respond(ctx, w, errNotFound, http.StatusNotFound)
	}

	if err != nil {
		return err
	}
	return web.Respond(ctx, w, present(result), http.StatusOK)
}

func present(val external.Account) validateResponse {
	return validateResponse{
		AccountName:   val.Name,
		AccountNumber: val.BankNumber,
		IsValid:       val.IsValid,
	}
}
