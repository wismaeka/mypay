package transaction

import (
	"net/http"

	"mypayment/business/web"
)

var (
	errBadRequest = &web.CustomError{
		Message:    "bad request",
		ErrorCode:  "BAD_REQUEST",
		StatusCode: http.StatusBadRequest,
	}
)
