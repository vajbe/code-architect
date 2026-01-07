package main

import (
	"sync"
	"time"
)

type Task struct {
	Id       int
	joinedAt time.Time
	res      chan string
}

type TaskStore struct {
	workQueue     []Task
	mu            *sync.RWMutex
	notifyQueue   chan bool
	responseQueue chan string
	stats         map[int]time.Duration
}

type StatsStore struct {
}
