package main

import (
	"log"

	"github.com/IBM/sarama"
)

func SendMessage(message string, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	producer.Input() <- msg

	return nil
}

func DeliveryHandler() {
	go func() {
		for {
			select {
			case msg := <-producer.Successes():
				log.Println("Delivered message to", msg.Topic)
			case err := <-producer.Errors():
				log.Println("Delivery failed:", err)
			}
		}
	}()
}

func CloseProducer() {
	producer.Close()
}
