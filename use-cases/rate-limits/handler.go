package main

import (
	"log"
	"net/http"
	"time"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("userId")
	log.Println("UserId: ", id)

	store.mu.Lock()
	defer store.mu.Unlock()
	if store.store[id] == nil {
		store.store[id] = &RateLimit{
			count:       1,
			windowStart: time.Now(),
		}
	} else {
		log.Print("Processing req")
		now := time.Now()
		if now.Sub(store.store[id].windowStart).Seconds() <= 1 && store.store[id].count >= MAX_LIMIT {
			log.Print("Too many requests for user", id)
			w.WriteHeader(http.StatusTooManyRequests)

			return
		} else if now.Sub(store.store[id].windowStart).Seconds() >= 1 {
			store.store[id].count = 1
			store.store[id].windowStart = time.Now()
		} else {
			store.store[id].count++
			// store.store[id].windowStart = time.Now()
		}
	}
	w.Write([]byte("Operation completed"))
}
