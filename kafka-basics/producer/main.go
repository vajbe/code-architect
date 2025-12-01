package main

import (
	"sync"
)

func main() {
	go Producer()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
