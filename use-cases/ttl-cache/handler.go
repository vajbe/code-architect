package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func NewCacheServer(store *CacheStore) *CacheServer {
	return &CacheServer{
		store: store,
	}
}

func (cs *CacheServer) SetCacheHandler(w http.ResponseWriter, r *http.Request) {

	typ := r.Method

	if typ == http.MethodPost {
		var requestBody SetCacheRequest
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&requestBody)
		if err != nil {
			res := &SetCacheResponse{
				Err:     err.Error(),
				Status:  http.StatusInternalServerError,
				Message: `Error while performing the operation`,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}
		defer r.Body.Close()

		ttl, err := time.ParseDuration(requestBody.ExpiresAt)
		if err != nil {
			res := &SetCacheResponse{
				Err:     err.Error(),
				Status:  http.StatusBadRequest,
				Message: `Invalid TTL format`,
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		cs.store.Set(requestBody.Key, requestBody.Value, ttl)

		res := &SetCacheResponse{
			Status:  http.StatusOK,
			Message: "Successfully key added",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (cs *CacheServer) GetCachehandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing 'key' query parameter"))
		return
	}
	value, ok := cs.store.Get(key)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key not found"))
		return
	}
	w.Write([]byte(value))
}
