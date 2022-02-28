package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/MijPeter/saxa/internal/error"
	"github.com/MijPeter/saxa/internal/service"
)

var routes = []route{
	newRoute(http.MethodGet, "/image", getImages),        // returns json list of image names
	newRoute(http.MethodPost, "/image", postImage),       // creates new image
	newRoute(http.MethodGet, "/image/([^/]+)", getImage), // returns image
}

func getImages(w http.ResponseWriter, r *http.Request) error {
	// some basic validation needed here
	// TODO

	images, _ := image.Query()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(images)
	return err
}

func postImage(w http.ResponseWriter, r *http.Request) error {
	// some basic validation needed here
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return err
	}

	in := r.MultipartForm.File
	if len(in) != 1 {
		return httperror.New("incorrect number of files", http.StatusBadRequest)
	}

	for name, headers := range in {
		x := headers[0]
		y, _ := x.Open()
		defer y.Close()
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, y); err != nil {
			return err;
		}

		image, err := image.Create(name, buf.Bytes())
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func getImage(w http.ResponseWriter, r *http.Request) error {
	// some basic validation needed here
	name := getParam(r, 0)

	image, err := image.Fetch(name)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(image)
	return err
}

func Controller(w http.ResponseWriter, r *http.Request) {
	router(w, r, routes)
}
