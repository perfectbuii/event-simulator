#!/bin/bash

# Function to run the Docker container with specified test
run_test() {
    local test_name=$1
    echo "Running test: $test_name"
    docker run --rm --network host f1-test-image run constant -r 1000/s -d 2s $test_name
}

# Check the argument and run the appropriate test
case $1 in
    franz-test-api)
        run_test testAPIWithFranzKafka
        ;;
    confluent-test-api)
        run_test testAPIWithConfluentKafka
        ;;
    franz-test-kafka)
        run_test testFranzKafkaProducer
        ;;
    confluent-test-kafka)
        run_test testConfluentKafkaProducer
        ;;
    *)
        echo "Usage: $0 {franz-test-api|confluent-test-api|franz-test-kafka|confluent-test-kafka}"
        exit 1
        ;;
esac