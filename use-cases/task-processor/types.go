package main

import "sync"

type TaskStore struct {
	mu    *sync.RWMutex
	tasks map[string]*Task
}

type Task struct {
	Id     string
	Status TaskStatus
}

type TaskStatus string
