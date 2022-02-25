package handler

import (
	"context"
	"net/http"
	"regexp"
)

func newRoute(method string, uri string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + uri + "$"), handler}
}

type route struct {
	method  string
	uri     *regexp.Regexp
	handler http.HandlerFunc
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

		r_.handler(w, r.WithContext(context.WithValue(r.Context(), paramsKey, args[1:])))
		return
	}

	if !allowed {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}
