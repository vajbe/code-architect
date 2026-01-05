package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Producer(id int, ch chan<- string, producerWg *sync.WaitGroup) {
	defer producerWg.Done()
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf(`log-%d-%d`, id, rand.Int())
		ch <- msg
		log.Println("Produced: \t\t", msg)
	}
}

func Consumer(id int, ch <-chan string, consumerWg *sync.WaitGroup, totalProcessedCount *atomic.Int32) {
	defer consumerWg.Done()
	for msg := range ch {
		time.Sleep(100 * time.Millisecond)
		log.Println("Processed: \t\t", msg, " \tby consumer ", id)
		totalProcessedCount.Add(1)
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	var ch chan string
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup
	var totalProcessedCount atomic.Int32

	ch = make(chan string, 5)
	for i := 1; i <= 2; i++ {
		producerWg.Add(1)
		go Producer(i, ch, &producerWg)
	}

	for i := 1; i <= 2; i++ {
		consumerWg.Add(1)
		go Consumer(i, ch, &consumerWg, &totalProcessedCount)
	}

	producerWg.Wait()
	close(ch)

	consumerWg.Wait()
	log.Println("\n\nTotal events processed:\t\t", totalProcessedCount.Load())
}
