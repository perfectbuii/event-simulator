franz-test-api:
	docker build -f test.Dockerfile -t test . && docker run --rm --network host test run constant -r 1000/s -d 2s testAPIWithFranzKafka

confluent-test-api:
	gdocker build -f test.Dockerfile -t test . && docker run --rm --network host test run constant -r 1000/s -d 2s testAPIWithConfluentKafka

franz-test-kafka:
	docker build -f test.Dockerfile -t test . && docker run --rm --network host test run constant -r 1000/s -d 2s testFranzKafkaProducer 

confluent-test-kafka:
	docker build -f test.Dockerfile -t test . && docker run --rm --network host test run constant -r 1000/s -d 2s testConfluentKafkaProducer

build:
	docker-compose down && docker-compose build --no-cache && docker-compose up

start:
	docker build -t myapp . && docker run --rm --network host myapp

it:
	docker build -f test.Dockerfile -t my-test-image . && docker run --rm --network host my-test-image