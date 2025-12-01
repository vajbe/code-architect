package main

import (
	"sync"
)

func main() {
	go Consumer()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
