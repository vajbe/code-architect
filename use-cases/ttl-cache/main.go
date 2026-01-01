package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	Addr string = ":8080"
)

var (
	server *http.Server
	mx     *http.ServeMux
)

func Initialize(cs *CacheStore) {
	mx = http.NewServeMux()
	cacheServer := NewCacheServer(cs)

	mx.HandleFunc("POST /api/set", cacheServer.SetCacheHandler)
	mx.HandleFunc("GET /api/get", cacheServer.GetCachehandler)
	server = &http.Server{
		Addr:    Addr,
		Handler: mx,
	}

	log.Println("Server started at", Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Server stopped unexpectedly Err: ", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	<-stop

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Server shutdown gracefully...")

}

func main() {
	cs := NewCacheStore()
	Initialize(cs)
	defer cs.Stop()
}
