package main

import (
	"context"
	"sync"
	"time"
)

func NewCacheStore() *CacheStore {
	ctx, cancel := context.WithCancel(context.Background())
	c := &CacheStore{
		data: make(map[string]*Entry),
		mu:   &sync.RWMutex{},
		ctx:  ctx, cancel: cancel,
	}
	go c.cleanUpJob()
	return c
}

func (cs *CacheStore) Set(key string, value string, ttl time.Duration) {
	cs.mu.Lock()

	cs.data[key] = &Entry{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}

	cs.mu.Unlock()
}

func (cs *CacheStore) Get(key string) (string, bool) {
	cs.mu.RLock()
	entry, ok := cs.data[key]
	if !ok {
		return "", false
	}
	cs.mu.RUnlock()

	if time.Now().After(entry.expiresAt) {
		cs.mu.Lock()
		delete(cs.data, key)
		cs.mu.Unlock()
		return "", false
	}
	return entry.value, true
}

func (cs *CacheStore) cleanUpJob() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			cs.mu.Lock()
			for k, v := range cs.data {
				if time.Now().After(v.expiresAt) {
					delete(cs.data, k)
				}
			}
			cs.mu.Unlock()
		case <-cs.ctx.Done():
			return
		}
	}

}

func (cs *CacheStore) Stop() {
	cs.cancel()
}
