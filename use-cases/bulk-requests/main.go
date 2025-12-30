package main

import (
	"log"
	"net/http"
)

func SendBulk() {
	// time.Sleep(5 * time.Second)

	client := &http.Client{
		// Timeout: 5 * time.Second,
	}

	for i := 0; i < 10; i++ {
		log.Println("Sending request ", i)
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/task",
			nil,
		)
		if err != nil {
			log.Println("request error:", err)
			continue
		}

		req.Header.Set("userId", "0002")

		resp, err := client.Do(req)
		if err != nil {
			log.Println("http error:", err)
			continue
		}

		// IMPORTANT
		resp.Body.Close()
	}
}

func main() {
	SendBulk()
}
