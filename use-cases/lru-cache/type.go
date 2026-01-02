package main

type CacheNode struct {
	key, value string
	prev, next *CacheNode
}

type LRUCache struct {
	cache    map[string]CacheNode
	head     *CacheNode
	tail     *CacheNode
	capacity int
}
