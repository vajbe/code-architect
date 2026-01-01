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
}
