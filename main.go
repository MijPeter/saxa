package main

import (
	"log"
	"net/http"

	"github.com/MijPeter/saxa/db"
	"github.com/MijPeter/saxa/internal/handler"
)

var serverPort = ":8080"

func main() {
	db.ConnectDb()
	defer db.DisconnectDb()
	log.Print("Connected to db!")
	http.HandleFunc("/", handler.Controller)
	log.Printf("Starting golang server on port %s", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
