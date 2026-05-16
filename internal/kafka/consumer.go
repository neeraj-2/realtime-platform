package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Start() {

	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{"localhost:9092"},
		"notification-group",
		config,
	)

	if err != nil {
		panic(err)
	}

	handler := ConsumerGroupHandler{}

	for {
		err := consumerGroup.Consume(
			context.Background(),
			[]string{"user-created"},
			handler,
		)

		if err != nil {
			log.Println(err)
		}
	}
}
