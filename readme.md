## Kafka Client Libraries Comparison

### Overview
This project uses two Go libraries for interacting with Kafka clusters: confluent-kafka-go and franz-go. Both libraries provide robust functionality for Kafka clients but differ in their features and performance characteristics.

### Libraries
**confluent-kafka-go** is a Go client for Kafka that is maintained by Confluent. It provides a high-level API for producing and consuming messages and is widely used in the Go community for its comprehensive feature set and reliability.

* High-performance producer and consumer
* Support for Kafka features like transactions, exactly-once semantics, and more
* Comprehensive documentation and community support

**franz-go**
franz-go is an idiomatic Go client for Kafka designed for performance and simplicity. It is a relatively new library but offers a range of features for producing and consuming Kafka messages efficiently.

* Efficient, high-throughput producer and consumer
* Built with Go idioms in mind, providing a more native Go experience
* Lightweight and performant

## Performance Testing with F1
I used F1 to conduct performance tests for both APIs and Kafka producers. The tests were performed under the following conditions:

* Payload Size: Identical payload sizes were used for both the APIs and Kafka producers to ensure consistency.
* Throughput: The tests simulated a load of 1000 requests per second.
* Duration: Each test was conducted over a duration of 2 seconds.

This setup allowed us to evaluate the performance and efficiency of each library under comparable conditions.

### Benchmark Results

**APIs**

* confluent-kafka-go API: api/v1/user-event
```bash
F1 Load Tester
Running testAPIWithConfluentKafka scenario for 2s at a rate of 1000/s constant rate, using distribution regular.

[   1s]  ✔  1000  ✘     0 (1000/s)   avg: 54.815027ms, min: 12.732776ms, max: 63.062858ms
[   2s]  Max Duration Elapsed - waiting for active tests to complete
[Teardown] ✔

Load Test Passed
2000 iterations started in 1.99124155s (1000/second)
Successful Iterations: 2000 (100.00%, 1000/second) avg: 54.031079ms, min: 10.733743ms, max: 63.062858ms
Full logs: /tmp/f1-testAPIWithConfluentKafka-789b-2024-08-01_21-18-14.log
```

* franz-go API: api/v1/order-event
```bash
F1 Load Tester
Running testAPIWithFranzKafka scenario for 2s at a rate of 1000/s constant rate, using distribution regular.

[   1s]  ✔  1000  ✘     0 (1000/s)   avg: 26.664394ms, min: 5.130325ms, max: 57.724944ms
[   2s]  Max Duration Elapsed - waiting for active tests to complete
[Teardown] ✔

Load Test Passed
2000 iterations started in 1.990963465s (1000/second)
Successful Iterations: 2000 (100.00%, 1000/second) avg: 24.802233ms, min: 4.1639ms, max: 57.724944ms
Full logs: /tmp/f1-testAPIWithFranzKafka-69bc-2024-08-01_21-17-49.log
```



**Kafka Producers**

* confluent-kafka-go Producer
```bash
F1 Load Tester
Running testConfluentKafkaProducer scenario for 2s at a rate of 1000/s constant rate, using distribution regular.

[   1s]  ✔  1000  ✘     0 (1000/s)   avg: 51.651617ms, min: 46.42722ms, max: 89.650031ms
[   2s]  Max Duration Elapsed - waiting for active tests to complete
[Teardown] ✔

Load Test Passed
2000 iterations started in 1.990825929s (1000/second)
Successful Iterations: 2000 (100.00%, 1000/second) avg: 49.643657ms, min: 46.351818ms, max: 89.650031ms
Full logs: /tmp/f1-testConfluentKafkaProducer-f3d5-2024-08-01_21-18-44.log
```

* franz-go Producer
```bash
F1 Load Tester
Running testFranzKafkaProducer scenario for 2s at a rate of 1000/s constant rate, using distribution regular.

[   1s]  ✔  1000  ✘     0 (1000/s)   avg: 4.505458ms, min: 954.733µs, max: 25.595585ms
[   2s]  Max Duration Elapsed - waiting for active tests to complete
[Teardown] ✔

Load Test Passed
2000 iterations started in 1.990659875s (1000/second)
Successful Iterations: 2000 (100.00%, 1000/second) avg: 3.252416ms, min: 804.419µs, max: 25.595585ms
Full logs: /tmp/f1-testFranzKafkaProducer-fc45-2024-08-01_21-19-14.log
```

## Usage 

```bash
# build project
make build
# run app
make start

# run test
make franz-test-api
make confluent-test-api
make franz-test-kafka
make confluent-test-kafka
```