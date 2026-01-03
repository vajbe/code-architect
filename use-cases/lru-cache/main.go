package main

import (
	"log"
	"net/http"
)

const (
	Addr string = ":8080"
)

func InitializeServer() {

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    Addr,
		Handler: mux,
	}

	cs := NewCacheServer()

	mux.HandleFunc("GET /api/get/{key}", cs.GetCacheHandler)
	mux.HandleFunc("POST /api/put", cs.PostCacheHandler)

	log.Println("Server started at ", Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("Error occurred while starting server; err:", err.Error())
	}

}
func main() {
	InitializeServer()
}
