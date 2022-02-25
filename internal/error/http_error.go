package err

import (
	"errors"
	"net/http"
)

// TODO https://blog.questionable.services/article/http-handler-error-handling-revisited/

type HttpError interface {
	Status() int
	Error() error
}

type httpError struct {
	status int
	err error
}

func (re *httpError) Status() int {
	return re.status
}

func (re *httpError) Err() error {
	return re.err
}

func New(errMsg string, status int) *HttpError {
	return &httpError{status, errors.New(errMsg)}
}

// errors
var IMAGE_NOT_FOUND = HttpError{
	http.StatusNotFound,
	errors.New("image has not been found")}

var IMAGE_NAME_EXISTS = HttpError{
	http.StatusConflict,
	errors.New("image with provided name already exists")}
