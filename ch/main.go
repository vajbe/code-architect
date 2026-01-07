package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const MAX_WORKERS int = 3

func (store *TaskStore) JoinIds(id1 Task, id2 Task) {
	id := rand.Int()
	res1 := fmt.Sprintf("%d First", id)
	res2 := fmt.Sprintf("%d Second", id)
	id1.res <- res1
	id2.res <- res2
	store.mu.Lock()
	store.stats[id] = time.Duration(id2.joinedAt.Sub(id1.joinedAt).Microseconds())
	store.mu.Unlock()
}

func Process(id int, store *TaskStore) {
	for range store.notifyQueue {
		log.Println("Received")
		store.mu.Lock()
		l := len(store.workQueue)
		if l >= 2 {
			task1 := store.workQueue[0]
			task2 := store.workQueue[1]
			// Remove elems from the queue
			store.workQueue = store.workQueue[2:]
			store.mu.Unlock()
			store.JoinIds(task1, task2)
		} else {
			store.mu.Unlock()
		}
	}
}

func (t *TaskStore) AddToQueue(task Task) {
	log.Println("In add to queue")
	t.mu.Lock()
	log.Println("Lock acquired")
	t.workQueue = append(t.workQueue, task)
	t.mu.Unlock()
	log.Println("Added to queue")
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		workQueue:   []Task{},
		mu:          &sync.RWMutex{},
		notifyQueue: make(chan bool),
		stats:       make(map[int]time.Duration),
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	store := NewTaskStore()
	for i := 0; i < MAX_WORKERS; i++ {
		go Process(i, store)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /joins", store.JoinHandler)
	mux.HandleFunc("GET /stats", store.StatsHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	//graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Shutdown gracefully...")

}

/* Program HTTP server which has two endpoints: POST /join This endpoint joins together pairs of clients calling it one after another.
It returns some random identifier of the pair + whether the client is the first or the second in the pair.
If a client is alone for 10 seconds (nobody else calls the endpoint), it returns "Timeout. No more connected clients."
GET /stats This endpoint returns identifiers and delays between the first and the second user in microseconds: {"id1": 50000, "id2": 10000}.
If there were more than 50'000 pairs, it only returns 50'000 with the shortest delay. Examples The first user makes POST /join. 2 seconds elapse, the query is blocked. The second user makes POST /join. The first gets "12342 First". The second gets "12342 Second". The user makes POST /join. 10 seconds elapse, the query is blocked. The user gets "Timeout. No more connected clients."
The user makes GET /stats. The server responds {"12342": 2000000}. */
