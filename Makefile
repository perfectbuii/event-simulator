init-infra:
	docker-compose down && docker-compose build --no-cache && docker-compose up

start-app:
	docker build -t myapp . && docker run --rm --network host myapp

build-test:
	docker build -f test.Dockerfile -t f1-test-image .

franz-test-api:
	./run-tests.sh franz-test-api

confluent-test-api:
	./run-tests.sh confluent-test-api

franz-test-kafka:
	./run-tests.sh franz-test-kafka 

confluent-test-kafka:
	./run-tests.sh confluent-test-kafka