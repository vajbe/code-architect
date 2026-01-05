package main

import (
	"log"
	"sync"
	"time"
)

const MAX_LIMIT = 10 * time.Second

func Pong(ping chan<- bool, pong <-chan bool, wg *sync.WaitGroup) {
	timer := time.NewTimer(MAX_LIMIT)
	defer wg.Done()
	defer timer.Stop()
	for {
		select {
		case <-pong:
			log.Println("Pong")
			time.Sleep(time.Second)
			ping <- true
		case <-timer.C:
			return
		}
	}
}

func Ping(ping <-chan bool, pong chan<- bool, wg *sync.WaitGroup) {
	timer := time.NewTimer(MAX_LIMIT)
	defer wg.Done()
	defer timer.Stop()
	for {
		select {
		case <-ping:
			log.Println("Ping")
			time.Sleep(time.Second)
			pong <- true
		case <-timer.C:
			return
		}
	}
}

func main() {

	ping := make(chan bool, 1)
	pong := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go Ping(ping, pong, &wg)
	go Pong(ping, pong, &wg)
	pong <- true
	wg.Wait()
	close(ping)
	close(pong)
	log.Println("Shutting down gracefully...")
}
