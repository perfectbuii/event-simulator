### [Event Simulator] - Franz-go Kafka Client Documentation

## Overview
This project utilizes the `franz-go` Kafka client library to interact with Kafka clusters. `franz-go` is a performant, idiomatic Go client for Kafka that supports high-throughput producers and consumers, transactions, and TLS encryption.

## Table of Contents
- Installation
- Configuration
  - Client Configuration
  - TLS Configuration
- Basic Usage
  - Producing Messages
  - Consuming Messages
- Securing with TLS
  - Generate Certificates
  - Configure TLS in franz-go
- Examples
  - Producing a Simple Message
  - Consuming Messages from a Topic
- Troubleshooting
  - Common Errors and Solutions
  - Debugging Tips
- Contributing
- License

## Installation
To start using `franz-go`, you need to install the package in your Go project. You can do this using `go get`.

```go 
go get github.com/twmb/franz-go/pkg/kgo
```


## Configuration

### Client Configuration
Create a Kafka client with your desired configuration options. Here's a basic example:

```go
package main

import (
    "github.com/twmb/franz-go/pkg/kgo"
    "log"
)

func main() {
    client, err := kgo.NewClient(
        kgo.SeedBrokers("localhost:9092"),
        kgo.ConsumerGroup("example-group"),
        kgo.ConsumeTopics("example-topic"),
    )
    if err != nil {
        log.Fatalf("Error creating Kafka client: %v", err)
    }
    defer client.Close()
}
```

### TLS Configuration
If your Kafka cluster is secured with TLS, you'll need to configure your client to use the appropriate certificates.

```go
package main

import (
    "crypto/tls"
    "crypto/x509"
    "log"
    "os"

    "github.com/twmb/franz-go/pkg/kgo"
)

func main() {
    client, err := kgo.NewClient(
        kgo.SeedBrokers("localhost:9092"),
        kgo.DialTLSConfig(&tls.Config{
            RootCAs: mustLoadCACerts("certs/ca-cert.pem"),
            Certificates: []tls.Certificate{mustLoadClientCert("certs/server-cert.pem", "certs/server-key.pem")},
        }),
    )
    if err != nil {
        log.Fatalf("Error creating Kafka client: %v", err)
    }
    defer client.Close()
}

func mustLoadCACerts(caPath string) *x509.CertPool {
    caCert, err := os.ReadFile(caPath)
    if err != nil {
        log.Fatalf("Failed to read CA certificate: %v", err)
    }
    pool := x509.NewCertPool()
    if !pool.AppendCertsFromPEM(caCert) {
        log.Fatalf("Failed to append CA certificate to pool")
    }
    return pool
}

func mustLoadClientCert(certPath, keyPath string) tls.Certificate {
    cert, err := tls.LoadX509KeyPair(certPath, keyPath)
    if err != nil {
        log.Fatalf("Failed to load client certificate: %v", err)
    }
    return cert
}

```


## Basic Usage

### Producing Messages
Producing messages to a Kafka topic is straightforward. Below is a simple example of how to produce a message using `franz-go`.

```go
package main

import (
    "context"
    "github.com/twmb/franz-go/pkg/kgo"
    "log"
)

func produceMessage(client *kgo.Client, topic string, key, value []byte) {
    record := &kgo.Record{
        Topic: topic,
        Key:   key,
        Value: value,
    }
    err := client.ProduceSync(context.Background(), record).FirstErr()
    if err != nil {
        log.Fatalf("Failed to produce message: %v", err)
    } else {
        log.Println("Message produced successfully")
    }
}

```


### Consuming Messages
Here's a simple example of consuming messages using `franz-go`.

```go
package main

import (
    "context"
    "log"
    "github.com/twmb/franz-go/pkg/kgo"
)

func consumeMessages(client *kgo.Client) {
    ctx := context.Background()

    for {
        fetches := client.PollFetches(ctx)
        fetches.EachPartition(func(p kgo.FetchTopicPartition) {
            for _, record := range p.Records {
                log.Printf("Received message from topic %s: key=%s value=%s", record.Topic, string(record.Key), string(record.Value))
            }
        })
    }
}

```

## Securing with TLS
### Generate Certificates
Here's a brief guide on generating certificates using OpenSSL.

```bash
# Create a new private key for the CA
openssl req -new -x509 -keyout ca-key.pem -out ca-cert.pem -days 365 -nodes -subj "/CN=MyCA"

# Generate a new private key for the server
openssl genrsa -out server-key.pem 2048

# Create a certificate signing request (CSR) for the server
openssl req -new -key server-key.pem -out server-csr.pem -subj "/CN=localhost"

# Sign the server certificate with the CA
openssl x509 -req -in server-csr.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -days 365

```
### Configure TLS in franz-go
Update your Kafka client configuration to use the certificates.

```go
package main

import (
    "crypto/tls"
    "crypto/x509"
    "log"
    "os"

    "github.com/twmb/franz-go/pkg/kgo"
)

func main() {
    client, err := kgo.NewClient(
        kgo.SeedBrokers("localhost:9092"),
        kgo.DialTLSConfig(&tls.Config{
            RootCAs: mustLoadCACerts("certs/ca-cert.pem"),
            Certificates: []tls.Certificate{mustLoadClientCert("certs/server-cert.pem", "certs/server-key.pem")},
        }),
    )
    if err != nil {
        log.Fatalf("Error creating Kafka client: %v", err)
    }
    defer client.Close()
}

func mustLoadCACerts(caPath string) *x509.CertPool {
    caCert, err := os.ReadFile(caPath)
    if err != nil {
        log.Fatalf("Failed to read CA certificate: %v", err)
    }
    pool := x509.NewCertPool()
    if !pool.AppendCertsFromPEM(caCert) {
        log.Fatalf("Failed to append CA certificate to pool")
    }
    return pool
}

func mustLoadClientCert(certPath, keyPath string) tls.Certificate {
    cert, err := tls.LoadX509KeyPair(certPath, keyPath)
    if err != nil {
        log.Fatalf("Failed to load client certificate: %v", err)
    }
    return cert
}

```

## Examples

**Producing a Simple Message**

```go 
package main

import (
    "context"
    "github.com/twmb/franz-go/pkg/kgo"
    "log"
)

func main() {
    client, err := kgo.NewClient(
        kgo.SeedBrokers("localhost:9092"),
    )
    if err != nil {
        log.Fatalf("Error creating Kafka client: %v", err)
    }
    defer client.Close()

    produceMessage(client, "example-topic", []byte("key"), []byte("value"))
}

func produceMessage(client *kgo.Client, topic string, key, value []byte) {
    record := &kgo.Record{
        Topic: topic,
        Key:   key,
        Value: value,
    }
    err := client.ProduceSync(context.Background(), record).FirstErr()
    if err != nil {
        log.Fatalf("Failed to produce message: %v", err)
    } else {
        log.Println("Message produced successfully")
    }
}

```

**Consuming Messages from a Topic**

```go
package main

import (
    "context"
    "log"
    "github.com/twmb/franz-go/pkg/kgo"
)

func main() {
    client, err := kgo.NewClient(
        kgo.SeedBrokers("localhost:9092"),
        kgo.ConsumerGroup("example-group"),
        kgo.ConsumeTopics("example-topic"),
    )
    if err != nil {
        log.Fatalf("Error creating Kafka client: %v", err)
    }
    defer client.Close()

    consumeMessages(client)
}

func consumeMessages(client *kgo.Client) {
    ctx := context.Background()

    for {
        fetches := client.PollFetches(ctx)
        fetches.EachPartition(func(p kgo.FetchTopicPartition) {
            for _, record := range p.Records {
                log.Printf("Received message from topic %s: key=%s value=%s", record.Topic, string(record.Key), string(record.Value))
            }
        })
    }
}

```


## Troubleshooting

### Common Errors and Solutions
- **Connection Errors:** Ensure your Kafka brokers are reachable from the client and the correct ports are open.
- **TLS Issues:** Verify that your certificates are correctly configured and the paths in your code match your file system.

### Debugging Tips
- Use `kgo` client logging options to get detailed logs and debug information.
- Ensure your Kafka broker is running and accessible.

## Contributing
Contributions are welcome! If you have suggestions for improvements or have found a bug, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
