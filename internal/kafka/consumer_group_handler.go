package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(
	sarama.ConsumerGroupSession,
) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(
	sarama.ConsumerGroupSession,
) error {
	return nil
}

func (ConsumerGroupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {

	for message := range claim.Messages() {

		log.Printf(
			"Received message: %s",
			string(message.Value),
		)

		session.MarkMessage(message, "")
	}

	return nil
}