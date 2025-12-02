package main

import (
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	producer *kafka.Producer
	topic    string
	err      error
	broker   string
)

func Initialize() {

	broker = "localhost:9092"
	topic = "raw_events"

	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})

	if err != nil {
		log.Fatal("Error starting kafka", err)
		return
	}

}

func main() {
	Initialize()
	DeliveryHandler()
	defer producer.Close()
	http.HandleFunc("/events", EventsHandler)
	log.Println("Server starting on :8080")
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		log.Fatal("Error starting server: ", er)
	}
}
