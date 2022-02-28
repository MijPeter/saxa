package handler

import (
	"net/http"

	httperror "github.com/MijPeter/saxa/internal/error"
)

func handleError(
	w http.ResponseWriter,
	r *http.Request,
	handler errorHandlerFunc) {

	err := handler(w, r)
	if err == nil {
		return
	}

	switch err := err.(type) {
	case *httperror.HttpError:
		http.Error(w, err.Error(), err.Status())
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
