package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func Consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"172.21.32.1:9092"}, config)
	if err != nil {
		log.Fatal("\nFailed to start Sarama consumer:", err)
	}
	defer consumer.Close()
	fmt.Print("\nConsumer started")

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)

	if err != nil {
		log.Fatal("\nFailed to start Sarama consumer:", err)
	}
	defer partitionConsumer.Close()

	batch := []string{}
	batchSize := 1000
	var batchStartOffset int64 = -1

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			if batchStartOffset == -1 {
				batchStartOffset = msg.Offset
			}
			batch = append(batch, string(msg.Value))

			if len(batch) == batchSize {
				endOffset := msg.Offset
				go ProcessBatch(batch, batchStartOffset, endOffset)
				batch = []string{}
				batchStartOffset = -1
			}

		case err := <-partitionConsumer.Errors():
			log.Printf("\nError: %s\n", err)
		}
	}

}

func ProcessBatch(batch []string, from, to int64) {
	fmt.Println("\n Processing batch ", from, " to ", to)
}
