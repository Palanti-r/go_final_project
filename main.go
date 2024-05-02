package main

import (
	"log"
	"net/http"
	"os"

	"github.com/palanti-r/go_final_project/db"
	"github.com/palanti-r/go_final_project/handlers"
)

const (
	webDir      = "./web"
	defaultPort = "7540"
)

func main() {
	db.InitDB()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = ":" + defaultPort
		log.Printf("[INFO] Set up port: %s\n", port)
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	http.HandleFunc("/api/nextdate", handlers.NextDateHandler)
	http.HandleFunc("/api/task", handlers.TaskHandler)
	http.HandleFunc("/api/task/done", handlers.TaskDoneHandler)
	http.HandleFunc("/api/tasks", handlers.TasksListHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("[ERROR] Listen And Serve: %s\n", err)
	}

}
