package handler

import (
	"context"
	"net/http"
	"regexp"
)

type errorHandlerFunc func(http.ResponseWriter, *http.Request) error

func newRoute(
	method string,
	uri string,
	handler errorHandlerFunc) route {
	return route{method, regexp.MustCompile("^" + uri + "$"), handler}
}

type route struct {
	method  string
	uri     *regexp.Regexp
	handler errorHandlerFunc
}

type key int

const (
	paramsKey key = iota
)

func getParam(r *http.Request, i int) string {
	params := r.Context().Value(paramsKey).([]string)
	if len(params) <= i {
		return ""
	}
	return params[i]
}

func router(w http.ResponseWriter, r *http.Request, routes []route) {
	var allowed bool
	uri := r.URL.Path
	for _, r_ := range routes {
		args := r_.uri.FindStringSubmatch(uri)

		if len(args) == 0 {
			continue
		}

		if r_.method != r.Method {
			allowed = false
			continue
		}

		handleError(
			w, 
			r.WithContext(context.WithValue(r.Context(), paramsKey, args[1:])), 
			r_.handler)

		return
	}

	if !allowed {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}
