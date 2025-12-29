package main

import (
	"log"
	"net/http"
	"time"
)

func SendBulk() {
	time.Sleep(5 * time.Second)

	for range 1000 {
		// log.Print(i)
		_, err := http.Get("http://localhost:8080/task")

		if err != nil {
			log.Fatal(err)
		}
	}
}
