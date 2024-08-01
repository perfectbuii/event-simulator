package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/linkedin/goavro/v2"
)

// SchemaRegistry represents a client for interacting with a Schema Registry.
type SchemaRegistry struct {
	url string
}

// NewSchemaRegistry creates a new SchemaRegistry client.
func NewSchemaRegistry(url string) *SchemaRegistry {
	return &SchemaRegistry{url: url}
}

// GetLatestSchema fetches the latest schema for a given subject (e.g., topic name).
func (sr *SchemaRegistry) GetLatestSchema(subject string) (*goavro.Codec, error) {
	// Construct the URL to fetch the latest schema version for the subject.
	url := fmt.Sprintf("%s/subjects/%s-value/versions/latest", sr.url, subject)

	// Make the HTTP request to fetch the schema.
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schema: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the response to get the schema definition.
	var schemaResponse struct {
		Schema string `json:"schema"`
	}

	if err := json.Unmarshal(body, &schemaResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal schema response: %w", err)
	}

	// Create a new Avro codec from the schema definition.
	codec, err := goavro.NewCodec(schemaResponse.Schema)
	if err != nil {
		return nil, fmt.Errorf("failed to create codec: %w", err)
	}

	return codec, nil
}

// SerializePayload serializes a payload using the Avro codec.
func (sr *SchemaRegistry) SerializePayload(codec *goavro.Codec, payload map[string]interface{}) ([]byte, error) {
	// Serialize the payload using the Avro codec.
	binary, err := codec.BinaryFromNative(nil, payload)
	if err != nil {
		log.Printf("Error serializing payload: %v", err)
		return nil, fmt.Errorf("failed to serialize payload: %w", err)
	}

	return binary, nil
}
