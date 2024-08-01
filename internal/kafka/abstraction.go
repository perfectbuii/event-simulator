package kafka

type KafkaHandler interface {
	ProduceMessage(value []byte, topic string) error
}
