package app_errors

import "errors"

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	ErrNotFound   = errors.New("Not found")
	ErrBadRequest = errors.New("Bad request")
)
