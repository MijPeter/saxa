package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	image "github.com/MijPeter/saxa/internal/service"
)

var routes = []route{
	newRoute(http.MethodGet, "/image", getImages),        // returns json list of image names
	newRoute(http.MethodPost, "/image", postImage),       // creates new image
	newRoute(http.MethodGet, "/image/([^/]+)", getImage), // returns image
}

func getImages(w http.ResponseWriter, r *http.Request) {
	// some basic validation needed here
	// TODO
	fmt.Println("getImages")

}

func postImage(w http.ResponseWriter, r *http.Request) {
	// some basic validation needed here
	// TODO

	fmt.Println("postImage")
}

func getImage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getImage")

	// get name of image from url
	name := "example"
	// some basic validation needed here

	image := image.Fetch(name)
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("debug")
	err := json.NewEncoder(w).Encode(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Controller(w http.ResponseWriter, r *http.Request) {
	controller(w, r, routes)
}
