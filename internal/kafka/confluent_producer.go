package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ConfluentKafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewConfluentKafkaProducer(brokers, topic string) (*ConfluentKafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        brokers,
		"security.protocol":        "ssl",
		"ssl.ca.location":          "./certs/ca.crt",
		"ssl.certificate.location": "./certs/broker.crt",
		"ssl.key.location":         "./certs/broker.key",
	})
	if err != nil {
		return nil, err
	}

	return &ConfluentKafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (kp *ConfluentKafkaProducer) ProduceMessage(value []byte, topic string) error {
	deliveryChan := make(chan kafka.Event)
	err := kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %v", m.TopicPartition.Error)
	}

	close(deliveryChan)

	return nil
}
