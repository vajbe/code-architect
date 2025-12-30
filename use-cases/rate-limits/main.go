package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type RateLimit struct {
	count       int
	windowStart time.Time
	mu          *sync.Mutex
}

type CacheStore struct {
	mu    *sync.Mutex
	store map[string]*RateLimit
}

var store *CacheStore

const MAX_LIMIT = 3

func NewCacheStore() *CacheStore {
	return &CacheStore{
		mu:    &sync.Mutex{},
		store: make(map[string]*RateLimit),
	}
}

func main() {
	mux := http.NewServeMux()
	store = NewCacheStore()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("/task", TaskHandler)

	go func() {
		log.Println("Http server started:8080")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Print("Server shutdown gracefully")

}

// - Allow N requests per second per user
// - Reject excess requests
// - Thread-safe implementation

// timestamp for last update
// current count
// userId
