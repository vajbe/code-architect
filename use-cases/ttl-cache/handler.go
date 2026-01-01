package main

import (
	"encoding/json"
	"net/http"
)

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

func SetCacheHandler(w http.ResponseWriter, r *http.Request) {

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

func GetCachehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get cache handler"))
}
