package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/form3tech-oss/f1/v2/pkg/f1"
	"github.com/form3tech-oss/f1/v2/pkg/f1/testing"
	"github.com/perfectbuii/event-simulator/internal/kafka"
)

func main() {
	f := f1.New()
	f.Add("testAPIWithConfluentKafka", testAPIWithConfluentKafka("http://localhost:8080/api/v1/user-event"))
	f.Add("testAPIWithFranzKafka", testAPIWithFranzKafka("http://localhost:8080/api/v1/order-event"))
	f.Add("testConfluentKafkaProducer", setupConfluentKafkaProducerTest("localhost:9092", "user-events"))
	f.Add("testFranzKafkaProducer", setupFranzKafkaProducerTest("localhost:9092", "order-events"))
	f.Execute()
}

func testAPIWithConfluentKafka(apiURL string) testing.ScenarioFn {
	return func(t *testing.T) testing.RunFn {
		// This is the run phase, executed for each iteration of the scenario.
		t.Logger().Infof("Setting up test scenario for API: %s", apiURL)

		// Cleanup logic can be added here if needed.
		t.Cleanup(func() {
			t.Logger().Infof("Cleaning up test scenario for API: %s", apiURL)
		})

		runFn := func(t *testing.T) {
			// This is the run phase, executed for each iteration of the scenario.

			// Prepare the payload.
			payload := map[string]interface{}{
				"user_id":    "12345",
				"event_type": "login",
				"timestamp":  time.Now().Format(time.RFC3339),
			}

			payloadBytes, err := json.Marshal(payload)
			t.Require().NoError(err)

			// Perform HTTP POST request to the API endpoint.
			res, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
			t.Require().NoError(err)
			defer res.Body.Close()

			// Verify the response.
			t.Require().Equal(http.StatusOK, res.StatusCode)

			t.Logger().Infof("API %s responded with status: %d", apiURL, res.StatusCode)
		}

		return runFn

	}
}

func testAPIWithFranzKafka(apiURL string) testing.ScenarioFn {
	return func(t *testing.T) testing.RunFn {
		// This is the run phase, executed for each iteration of the scenario.
		t.Logger().Infof("Setting up test scenario for API: %s", apiURL)

		// Cleanup logic can be added here if needed.
		t.Cleanup(func() {
			t.Logger().Infof("Cleaning up test scenario for API: %s", apiURL)
		})

		runFn := func(t *testing.T) {
			// This is the run phase, executed for each iteration of the scenario.

			// Prepare the payload.
			payload := map[string]interface{}{
				"order_id":   "12345",
				"event_type": "login",
				"timestamp":  time.Now().Format(time.RFC3339),
			}

			payloadBytes, err := json.Marshal(payload)
			t.Require().NoError(err)

			// Perform HTTP POST request to the API endpoint.
			res, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
			t.Require().NoError(err)
			defer res.Body.Close()

			// Verify the response.
			t.Require().Equal(http.StatusOK, res.StatusCode)

			t.Logger().Infof("API %s responded with status: %d", apiURL, res.StatusCode)
		}

		return runFn

	}
}

// Setup test scenario for Confluent Kafka Producer
func setupConfluentKafkaProducerTest(broker, topic string) testing.ScenarioFn {
	// Create a Confluent Kafka producer
	producer, err := kafka.NewConfluentKafkaProducer(broker, topic)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s\n", err)
	}
	return func(t *testing.T) testing.RunFn {

		// Run function executed for each iteration
		return func(t *testing.T) {
			// Produce a message
			payload := map[string]interface{}{
				"user_id":    "12345",
				"event_type": "login",
				"timestamp":  time.Now().Format(time.RFC3339),
			}

			payloadBytes, err := json.Marshal(payload)
			t.Require().NoError(err)

			err = producer.ProduceMessage(payloadBytes, topic)
			t.Require().NoError(err)

			t.Logger().Infof("Confluent Kafka message sent successfully")
		}
	}
}

// Setup test scenario for Franz Kafka Producer
func setupFranzKafkaProducerTest(broker, topic string) testing.ScenarioFn {
	// Create a Franz Kafka client
	franzProducer, err := kafka.NewFranzKafkaProducer(broker, topic)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s\n", err)
	}
	return func(t *testing.T) testing.RunFn {
		// Run function executed for each iteration
		return func(t *testing.T) {
			// Produce a message
			payload := map[string]interface{}{
				"order_id":   "12345",
				"event_type": "login",
				"timestamp":  time.Now().Format(time.RFC3339),
			}

			payloadBytes, err := json.Marshal(payload)
			t.Require().NoError(err)

			err = franzProducer.ProduceMessage(payloadBytes, topic)
			t.Require().NoError(err)

			t.Logger().Infof("Franz Kafka message sent successfully")
		}
	}
}
