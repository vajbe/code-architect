package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"172.21.32.1:9092"}, config)
	if err != nil {
		log.Fatal("\nFailed to start Sarama producer:", err)
	}
	defer producer.Close()
	fmt.Print("\nProducer started")
	for {

		message := sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.StringEncoder("Hello World! " + time.Now().Format("2006-01-02 15:04:05")),
		}

		_, _, err := producer.SendMessage(&message)
		if err != nil {
			log.Fatal("\nFailed to send message:", err)
		}

		// time.Sleep(2 * time.Second)
		fmt.Print("\nMessage sent")

	}

}
