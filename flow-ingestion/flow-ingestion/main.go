package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	consumer sarama.ConsumerGroup
	topic    string
	err      error
	broker   string
	es       *elasticsearch.TypedClient
)

func Initialize() {

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	broker = "localhost:9092"
	topic = "raw_events"

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err = sarama.NewConsumerGroup([]string{broker}, "flow-ingestion-group", config)

	if err != nil {
		log.Fatal("Error starting kafka", err)
		return
	}

	es, _ = elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

}

func main() {
	Initialize()
	DeliveryHandler()

	// defer es.Close()
}
