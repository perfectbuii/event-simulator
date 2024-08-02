build:
	docker-compose down && docker-compose build --no-cache && docker-compose up -d && docker build -t myapp . && docker build -f test.Dockerfile -t f1-test-image .

start-app:
	docker run --rm --network host myapp

franz-test-api:
	./run-tests.sh franz-test-api

confluent-test-api:
	./run-tests.sh confluent-test-api

franz-test-kafka:
	./run-tests.sh franz-test-kafka 

confluent-test-kafka:
	./run-tests.sh confluent-test-kafka