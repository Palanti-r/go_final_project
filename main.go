package main

import (
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
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}

}
