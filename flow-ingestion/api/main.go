package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/events", EventsHandler)
	log.Println("Server starting on :8080")
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		log.Fatal("Error starting server: ", er)
	}
}
