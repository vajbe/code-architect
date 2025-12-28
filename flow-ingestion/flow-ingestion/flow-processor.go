package main

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

func IndexDocuments(batch []sarama.ConsumerMessage) {
	for _, document := range batch {
		res, _ := es.Index("flow-events").
			Request(document).
			Do(context.Background())

		if res == nil {
			log.Println("Error indexing document", document)
			continue
		}
	}
}

func FlowProcessor(batch []sarama.ConsumerMessage) {

}
