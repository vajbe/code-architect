package main

import (
	"fmt"
	"log"
	"sync"
)

type SingletonObject struct {
}

var syncOnce sync.Once
var singletonObject *SingletonObject
var mu = sync.Mutex{}

func getSingletonObject() *SingletonObject {
	if singletonObject == nil {
		mu.Lock()
		defer mu.Unlock()
		if singletonObject == nil {
			syncOnce.Do(func() {
				log.Println("Creating singleton object...")
				singletonObject = &SingletonObject{}
			})
		} else {
			log.Println("Object already exists")
		}

	} else {
		log.Println("Object already created")
	}
	return singletonObject
}

func main() {
	for i := 0; i < 20; i++ {
		go getSingletonObject()
	}
	fmt.Scanln()
}
