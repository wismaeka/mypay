// Package web includes all data that is used in the business layer
package web

// CustomError is the defined error
type CustomError struct {
	Message    string `json:"message"`
	ErrorCode  string `json:"error_code"`
	StatusCode int    `json:"-"`
}

// Error makes struct CustomError implement the error interface
func (e *CustomError) Error() string {
	return e.Message
}
