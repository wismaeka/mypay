package web

import (
	"encoding/json"
	"net/http"
)

// Decode reads the body of an HTTP request looking for a JSON document.
// The body is decoded into the provided value, which must be writable.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}
