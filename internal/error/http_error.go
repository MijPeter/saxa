package httperror

import (
	"errors"
	"net/http"
)

type HttpError struct {
	status int
	err    error
}

func (he *HttpError) Error() string {
	return he.err.Error()
}

func (he *HttpError) Status() int {
	return he.status
}

func New(errMsg string, status int) error {
	return &HttpError{status, errors.New(errMsg)}
}

// errors
var IMAGE_NOT_FOUND = &HttpError{
	http.StatusNotFound,
	errors.New("image has not been found")}

var IMAGE_NAME_EXISTS = &HttpError{
	http.StatusConflict,
	errors.New("image with provided name already exists")}
