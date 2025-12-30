package main

import (
	"log"
	"net/http"
)

func TaskHandler(store *CacheStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("userId")
		if id == "" {
			id = r.RemoteAddr
		}

		isAllowed := isRequestAllowed(id, store)
		if !isAllowed {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit exceeded"))
			log.Println("Flow dropped")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Request accepted"))
		log.Printf("Flow processed")
	}
}
