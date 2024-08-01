package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

// FranzKafkaProducer is an implementation of KafkaHandler using franz-go.
type FranzKafkaProducer struct {
	client *kgo.Client
}

// NewFranzKafkaProducer creates a new FranzKafkaProducer.
func NewFranzKafkaProducer(brokers string, topic string) (*FranzKafkaProducer, error) {
	// Load CA certificate
	caCert, err := ioutil.ReadFile("./certs/ca.crt")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}

	// Create a certificate pool from CA
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Load client certificate
	cert, err := tls.LoadX509KeyPair("./certs/broker.crt", "./certs/broker.key")
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	// Set up TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers),
		kgo.ConsumeTopics(topic), // Topics can be configured as needed
		kgo.DialTLSConfig(tlsConfig),
	)
	if err != nil {
		return nil, err
	}

	return &FranzKafkaProducer{client: client}, nil
}

// ProduceMessage produces a message to the specified Kafka topic.
func (fp *FranzKafkaProducer) ProduceMessage(value []byte, topic string) error {
	record := &kgo.Record{
		Topic: topic,
		Value: value,
	}

	if err := fp.client.ProduceSync(context.Background(), record).FirstErr(); err != nil {
		return err
	}

	return nil
}
