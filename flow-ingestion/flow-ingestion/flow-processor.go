package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func IndexDocuments(batch []sarama.ConsumerMessage) {
	for _, document := range batch {
		var doc map[string]interface{}
		json.Unmarshal(document.Value, &doc)

		res, err := es.Index("flow-events").
			Request(doc).
			Do(context.Background())

		if res == nil {
			log.Println("Error indexing document", err)
			continue
		}
	}
}

func FlowProcessor(batch []sarama.ConsumerMessage) {
	IndexDocuments(batch)
}
