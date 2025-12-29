package main

import (
	"net/http"
	"time"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	id := time.Now().Format("150405.000")
	store.mu.Lock()
	store.tasks[id] = &Task{
		Id:     id,
		Status: PENDING,
	}
	store.mu.Unlock()
	taskQueue <- id
	w.Write([]byte("Task Created: " + id))
	// w.WriteHeader(http.StatusOK)
}
