#!/bin/bash

# Load environment variables
source /topics.env

# Wait for Kafka to be ready
echo "Waiting for Kafka to be ready..."
sleep 10

# Create topics using rpk
rpk topic create $USER_EVENTS_TOPIC -p 1 -r 1 --brokers=$KAFKA_BROKERS
rpk topic create $ORDER_EVENTS_TOPIC -p 1 -r 1 --brokers=$KAFKA_BROKERS

echo "Topics created successfully."
