package main

import (
	"github.com/palanti-r/go_final_project/db"
	"log"
	"net/http"
	"os"
)

const (
	webDir      = "./web"
	defaultPort = "7540"
)

func main() {
	db.InitlDB()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = ":" + defaultPort
		log.Printf("[INFO] Set up port: %s\n", port)
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalf("[ERROR] Listen And Serve: %s\n", err)
	}

}
