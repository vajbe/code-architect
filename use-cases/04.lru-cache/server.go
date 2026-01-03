package main

import (
	"encoding/json"
	"net/http"
)

type CacheServer struct {
	store *LRUCache
}

type CacheGetResponse struct {
	Message string `json:"msg"`
	Value   string `json:"value,omitempty"`
	Err     string `json:"error,omitempty"`
}

type CacheSetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewCacheServer() *CacheServer {
	return &CacheServer{
		store: NewLRUCache(10),
	}
}

func (cs *CacheServer) GetCacheHandler(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	v, ok := cs.store.Get(key)

	if ok {
		res := &CacheGetResponse{
			Message: `Success`,
			Value:   v,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := &CacheGetResponse{
		Message: `Key not found`,
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)

}

func (cs *CacheServer) PostCacheHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var req CacheSetRequest
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cs.store.Put(req.Key, req.Value)

	res := &CacheGetResponse{
		Message: `Success`,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
