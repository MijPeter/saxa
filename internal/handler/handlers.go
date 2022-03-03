package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MijPeter/saxa/internal/handler/util"
	image "github.com/MijPeter/saxa/internal/service"
)

var routes = []route{
	newRoute(http.MethodOptions, "*", options),
	newRoute(http.MethodGet, "/image", getImages),        // returns json list of image names
	newRoute(http.MethodPost, "/image", postImage),       // creates new image
	newRoute(http.MethodGet, "/image/([^/]+)", getImage), // returns image
}

func getImages(w http.ResponseWriter, r *http.Request) error {
	images, err := image.Query()

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(images)
}

func postImage(w http.ResponseWriter, r *http.Request) error {
	name, bytes, err := util.ParseFile(r)
	if err != nil {
		return err
	}

	image, err := image.Create(name, bytes)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(image)
}

func getImage(w http.ResponseWriter, r *http.Request) error {
	name := getParam(r, 0)

	image, err := image.Fetch(name)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/image")
	_, err = w.Write(image.Content)
	return err
}

func options(w http.ResponseWriter, r *http.Request) error {
	// Depending on the Method different kind of Headers can be set

	switch method := r.Header.Get("Access-Control-Request-Method"); method {
	case http.MethodPost:
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
	return nil
}

func Controller(w http.ResponseWriter, r *http.Request) {
	// this header is only for running simple app from localhost
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	router(w, r, routes)
}
