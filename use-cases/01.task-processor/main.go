package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	PENDING   = "PENDING"
	COMPLETED = "COMPLETED"
	RUNNING   = "RUNNING"
)

var (
	store          *TaskStore
	taskQueue      chan string
	tasksProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tasks_processed_total",
		Help: "The total number of processed tasks",
	})
)

const MAX_WORKERS = 3

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[string]*Task),
	}
}

func (s *TaskStore) Update(id string, status TaskStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[id].Status = status
}

func Work(id int, que chan string, s *TaskStore) {
	for taskId := range que {
		s.Update(taskId, RUNNING)
		time.Sleep(2 * time.Second)
		s.Update(taskId, COMPLETED)
		tasksProcessed.Inc()
		fmt.Printf("TaskId: %s completed by worker WorkerId: %d\n", taskId, id)
	}
}

func main() {
	taskQueue = make(chan string, 10)
	store = NewTaskStore()

	for i := 0; i < MAX_WORKERS; i++ {
		go Work(i, taskQueue, store)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/task", TasksHandler)
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	go func() {
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go SendBulk()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}
