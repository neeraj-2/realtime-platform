package kafka

import (
	"encoding/json"


	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer() *KafkaProducer {

	configKafka := sarama.NewConfig()

	configKafka.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(
		[]string{"localhost:9092"},
		configKafka,
	)

	if err != nil {
		panic(err)
	}

	return &KafkaProducer{
		Producer: producer,
	}
}

func (k *KafkaProducer) Publish(
	topic string,
	payload interface{},
) error {

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonPayload),
	}

	_, _, err = k.Producer.SendMessage(message)

	return err
}