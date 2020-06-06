package errors

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// New returns a new instance of *AppError.
func New(code int, options ...Option) *AppError {
	appErr := AppError{Code: code}
	for _, option := range options {
		option(&appErr)
	}
	return &appErr
}

// AppError represents application business errors.
type AppError struct {
	Code             int           `json:"code"`
	Message          string        `json:"message"`
	ValidationErrors []AppError    `json:"validation_errors,omitempty"`
	OriginalError    error         `json:"-"`
	Params           []interface{} `json:"-"`
}

// Error returns AppError as a string.
func (_this *AppError) Error() string {
	data, _ := json.Marshal(_this)
	return string(data)
}

type Option func(appError *AppError)

func WithMessage(message string) Option {
	return func(appError *AppError) {
		appError.Message = message
	}
	ioutil.NopCloser()
}
