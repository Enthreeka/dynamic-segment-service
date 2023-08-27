package apperror

import (
	"errors"
	"fmt"
)

var (
	ErrInernalServerError = NewAppError(errors.New("internal server error"), "invalid request body")
)

var (
	ErrSegmentsNotFound = NewAppError(errors.New("segments_not_found"), "no segments")
)

type statusMessage string

type AppError struct {
	Err error         `json:"-"`
	Msg statusMessage `json:"message"`
}

func (a *AppError) Error() string {
	return fmt.Sprintf("%s: %s", a.Msg, a.Err)
}

func (a *AppError) Unwrap() error {
	return a.Err
}

func NewAppError(err error, msg statusMessage) *AppError {
	return &AppError{
		Err: err,
		Msg: msg,
	}
}
