package controller

import (
	"net/http"
	"regexp"
)

func newRoute(method string, uri string, handler handlerFunc) route {
	return route{method, regexp.MustCompile("^" + uri + "$"), handler}
}

type route struct {
	method  string
	uri     *regexp.Regexp
	handler handlerFunc
}

type handlerFunc func (http.ResponseWriter, *http.Request, ...string)

func controller(w http.ResponseWriter, r *http.Request, routes []route) {
	var allowed bool
	uri := []byte(r.URL.Path)
	for _, r_ := range routes {
		if !r_.uri.Match(uri) {
			continue
		}
		if r_.method != r.Method {
			allowed = false
			continue
		}

		r_.handler(w, r, r_.args)
		return
	}

	if !allowed {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}
