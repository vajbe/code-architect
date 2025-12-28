package main

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// ConsumerGroupHandler represents a Sarama consumer group handler
type ConsumerGroupHandler struct{}

const BatchSize = 100

// Setup is run at the beginning of a new session, before ConsumeClaim
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	log.Println("ConsumerGroupHandler.Setup: Consumer group session started")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	log.Println("ConsumerGroupHandler.Cleanup: Consumer group session ended")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// log.Printf("ConsumerGroupHandler.ConsumeClaim: Starting consumer loop for topic %q partition %d", claim.Topic(), claim.Partition())
	batch := []sarama.ConsumerMessage{}
	for message := range claim.Messages() {
		// log.Printf("Message topic:%q partition:%d offset:%d\n", message.Topic, message.Partition, message.Offset)

		batch = append(batch, *message)
		session.MarkMessage(message, "")
		if len(batch) == BatchSize {
			log.Println("ConsumerGroupHandler.ConsumeClaim: Processing batch of size", BatchSize, len(batch))
			FlowProcessor(batch)
			batch = []sarama.ConsumerMessage{}
		}
	}
	// log.Printf("ConsumerGroupHandler.ConsumeClaim: Channel closed for topic %q partition %d", claim.Topic(), claim.Partition())
	return nil
}

func DeliveryHandler() {
	groupHandler := ConsumerGroupHandler{}
	ctx := context.Background()

	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := consumer.Consume(ctx, []string{topic}, groupHandler); err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
	}
}
