package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/perfectbuii/event-simulator/internal/kafka"
	"github.com/perfectbuii/event-simulator/internal/schema"
)

type APIHandler struct {
	kafkaProducer  kafka.KafkaHandler
	schemaRegistry *schema.SchemaRegistry
	topic          string
}

type Payload struct {
	Key   string                 `json:"key"`
	Value map[string]interface{} `json:"value"`
}

func NewAPIHandler(kp kafka.KafkaHandler, sr *schema.SchemaRegistry, topic string) *APIHandler {
	return &APIHandler{
		kafkaProducer:  kp,
		schemaRegistry: sr,
		topic:          topic,
	}
}

func (h *APIHandler) ProduceMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch the latest schema for the topic
	codec, err := h.schemaRegistry.GetLatestSchema(h.topic)
	if err != nil {
		http.Error(w, "Failed to retrieve schema", http.StatusInternalServerError)
		return
	}

	// Validate and serialize the payload
	serializedData, err := h.schemaRegistry.SerializePayload(codec, msg)
	if err != nil {
		http.Error(w, "Payload serialization failed", http.StatusBadRequest)
		return
	}

	// Produce the message to Kafka
	if err := h.kafkaProducer.ProduceMessage(serializedData, h.topic); err != nil {
		http.Error(w, "Failed to produce message", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message produced successfully")
}
