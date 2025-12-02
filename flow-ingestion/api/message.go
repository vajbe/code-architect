package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SendMessage(message string, topic string) error {
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryHandler() {
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Println("Delivery failed: ", ev.TopicPartition.Error)
				} else {
					log.Println("Delivered message to ", ev.TopicPartition)
				}
			}
		}
	}()
}

func CloseProducer() {
	producer.Flush(10 * 1000)
	producer.Close()
}
