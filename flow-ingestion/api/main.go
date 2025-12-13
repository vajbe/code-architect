package main

import (
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

var (
	producer sarama.AsyncProducer
	topic    string
	err      error
	broker   string
)

func Initialize() {

	broker = "localhost:9092"
	topic = "raw_events"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err = sarama.NewAsyncProducer([]string{broker}, config)

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
