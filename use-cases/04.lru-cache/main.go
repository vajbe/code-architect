package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	Addr string = ":8080"
)

func InitializeAndStartServer() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    Addr,
		Handler: mux,
	}

	cs := NewCacheServer()

	mux.HandleFunc("GET /api/get/{key}", cs.GetCacheHandler)
	mux.HandleFunc("POST /api/put", cs.PostCacheHandler)

	log.Println("Server started at ", Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Error occurred while starting server; err:", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Server shutdown gracefully")
}
func main() {
	InitializeAndStartServer()
}
