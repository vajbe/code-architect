package main

import (
	"log"
	"net/http"
)

const (
	Addr string = ":8080"
)

var (
	server *http.Server
)

func Initialize() {
	mx := http.NewServeMux()

	mx.HandleFunc("POST /api/set", SetCacheHandler)
	mx.HandleFunc("GET /api/get", GetCachehandler)
	server = &http.Server{
		Addr:    ":8080",
		Handler: mx,
	}

	log.Println("Server started at", Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Server stopped unexpectedly Err: ", err.Error())
	}
}

func main() {
	// cs := NewCacheStore()
	// cs.data["Vivek"]
	Initialize()
}
