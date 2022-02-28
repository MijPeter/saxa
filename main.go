package main

import (
	"log"
	"net/http"

	"github.com/MijPeter/saxa/internal/handler"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.Controller)
	log.Printf("Starting golang server on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
