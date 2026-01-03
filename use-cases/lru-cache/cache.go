package main

const MAX_CACHE_LIMIT = 2

func NewLRUCache(capacity int) *LRUCache {
	head := &CacheNode{}
	tail := &CacheNode{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		cache:    make(map[string]*CacheNode),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (lr *LRUCache) addToFront(entryNode *CacheNode) {
	entryNode.next = lr.head.next
	entryNode.prev = lr.head
	lr.head.next.prev = entryNode
	lr.head.next = entryNode
}

func (lr *LRUCache) remove(node *CacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lr *LRUCache) Get(key string) (string, bool) {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	data, ok := lr.cache[key]

	if ok {
		lr.remove(data)
		lr.addToFront(data)
		return data.value, true
	} else {
		return "", false
	}
}

func (lr *LRUCache) Put(key string, value string) {

	// check if exists
	lr.mu.Lock()
	cacheEntry, ok := lr.cache[key]
	defer lr.mu.Unlock()

	if ok {
		cacheEntry.value = value

		lr.remove(cacheEntry)
		lr.addToFront(cacheEntry)
	} else {
		// check if full capacity
		if len(lr.cache) == lr.capacity {
			lru := lr.tail.prev
			lr.remove(lru)
			delete(lr.cache, lru.key)
		}

		newNode := &CacheNode{
			key:   key,
			value: value,
		}
		lr.addToFront(newNode)
		lr.cache[key] = newNode
	}

}
