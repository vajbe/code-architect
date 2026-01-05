package main

import (
	"log"
	"sync"
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

const MAX_LIMIT = 3

func NewCacheStore() *CacheStore {
	return &CacheStore{
		mu:    &sync.Mutex{},
		store: make(map[string]*RateLimit),
	}
}

func (cs *CacheStore) getLimiter(userId string) *RateLimit {

	cs.mu.Lock()
	defer cs.mu.Unlock()

	if limiter, exists := cs.store[userId]; exists {
		return limiter
	}

	limiter := &RateLimit{
		count:       0,
		windowStart: time.Now(),
		mu:          &sync.Mutex{},
	}

	cs.store[userId] = limiter
	return limiter
}

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
