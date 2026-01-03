package main

import (
	"log"
	"time"
)

func isRequestAllowed(id string, store *CacheStore) bool {
	limiter := store.getLimiter(id)
	now := time.Now()

	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	if now.Sub(limiter.windowStart) >= time.Second {
		limiter.count = 0
		limiter.windowStart = now
		log.Println("Limits reset")
	}

	if limiter.count > MAX_LIMIT-1 {
		return false
	}

	limiter.count++
	return true
}
