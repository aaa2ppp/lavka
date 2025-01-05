package helper

// type HTTPError struct {
// 	StatusCode int
// 	Message    string
// }

// func (e *HTTPError) Error() string {
// 	return fmt.Sprintf("%d %s", e.StatusCode, e.Message)
// }

// var (
// 	ErrBadRequest    = &HTTPError{400, "bad request"}
// 	ErrNotFound      = &HTTPError{404, "not found"}
// 	ErrInternalError = &HTTPError{500, "internal error"}
// )

import (
	"lavka/internal/model"
)

type HTTPError = model.HTTPError

var (
	ErrBadRequest    = model.ErrBadRequest
	ErrNotFound      = model.ErrNotFound
	ErrInternalError = model.ErrInternalError
)
