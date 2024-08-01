franz-test-api:
	go run ./f1/main.go run constant -r 1000/s -d 2s testAPIWithFranzKafka

confluent-test-api:
	go run ./f1/main.go run constant -r 1000/s -d 2s testAPIWithConfluentKafka

franz-test-kafka:
	go run ./f1/main.go run constant -r 1000/s -d 2s testFranzKafkaProducer 

confluent-test-kafka:
	go run ./f1/main.go run constant -r 1000/s -d 2s testConfluentKafkaProducer

build:
	docker-compose down && docker-compose build --no-cache && docker-compose up

start:
	go run cmd/main.go