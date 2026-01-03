package main

import (
	"context"
	"sync"
	"time"
)

type Entry struct {
	value     string
	expiresAt time.Time
}

type CacheStore struct {
	data   map[string]*Entry
	mu     *sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

type SetCacheRequest struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ExpiresAt string `json:"ttl"`
}

type SetCacheResponse struct {
	Err     string `json:"err,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CacheServer struct {
	store *CacheStore
}
