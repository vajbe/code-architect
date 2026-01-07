package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (store *TaskStore) JoinHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Id:       rand.Int(),
		joinedAt: time.Now(),
		res:      make(chan string),
	}
	log.Println("Adding to queue")
	store.AddToQueue(task)
	store.notifyQueue <- true
	select {
	case <-time.After(10 * time.Second):
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Request timed out"))
	case r := <-task.res:
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(r))
	}
}

func (store *TaskStore) StatsHandler(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(store.stats)
}
