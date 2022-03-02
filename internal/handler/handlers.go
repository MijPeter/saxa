package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MijPeter/saxa/db"
	"github.com/MijPeter/saxa/internal/handler/util"
	"github.com/MijPeter/saxa/internal/service"
)

var routes = []route{
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

func Controller(w http.ResponseWriter, r *http.Request) {
	if db.DB == nil {
		fmt.Printf("ASDF")
	}
	router(w, r, routes)
}
