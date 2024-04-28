package main

import (
	"log"
	"net/http"
	"os"
)

const (
	webDir      = "./web"
	defaultPort = "7540"
)

func main() {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = ":" + defaultPort
		log.Printf("Set up port %s\n", port)
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalf("Listen And Serve - ", err)
	}

}
