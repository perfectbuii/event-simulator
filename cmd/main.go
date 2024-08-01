package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/perfectbuii/event-simulator/internal/api"
	"github.com/perfectbuii/event-simulator/internal/kafka"
	"github.com/perfectbuii/event-simulator/internal/schema"
)

func main() {
	// Default values for environment variables
	kafkaBrokers := getEnv("KAFKA_BROKERS", "localhost:9092")
	schemaRegistryURL := getEnv("SCHEMA_REGISTRY_URL", "http://localhost:8081")
	userEventsTopic := getEnv("USER_EVENTS_TOPIC", "user-events")
	orderEventsTopic := getEnv("ORDER_EVENTS_TOPIC", "order-events")

	confluentProducer, err := kafka.NewConfluentKafkaProducer(kafkaBrokers, userEventsTopic)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s\n", err)
	}

	franzProducer, err := kafka.NewFranzKafkaProducer(kafkaBrokers, "")
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s\n", err)
	}

	// Initialize Schema Registry client
	schemaRegistry := schema.NewSchemaRegistry(schemaRegistryURL)

	// Initialize API handlers for each event type
	userEventHandler := api.NewAPIHandler(confluentProducer, schemaRegistry, userEventsTopic)
	orderEventHandler := api.NewAPIHandler(franzProducer, schemaRegistry, orderEventsTopic)

	// Setup HTTP router
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/user-event", userEventHandler.ProduceMessageHandler).Methods("POST")
	router.HandleFunc("/api/v1/order-event", orderEventHandler.ProduceMessageHandler).Methods("POST")

	// Start HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
