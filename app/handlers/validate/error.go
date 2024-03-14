package validate

import (
	"net/http"

	"mypayment/business/web"
)

var (
	errNotFound = &web.CustomError{
		Message:    "no data found",
		ErrorCode:  "NOT_FOUND",
		StatusCode: http.StatusNotFound,
	}
)
