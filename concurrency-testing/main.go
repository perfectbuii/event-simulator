package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	const concurrencyLevel = 2000 // Set the number of concurrent goroutines

	var (
		count int            // Counter for successful requests
		wg    sync.WaitGroup // WaitGroup to wait for all goroutines to complete
		mutex sync.Mutex     // Mutex to protect the count variable
	)

	stopChan := make(chan struct{}) // Channel to signal when the 2-second duration is over

	// Launch a goroutine to close the channel after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		close(stopChan)
	}()

	// Channel to control concurrency
	concurrencyChan := make(chan struct{}, concurrencyLevel)

	// Function to perform HTTP request
	request := func() {
		defer wg.Done()

		// Prepare payload
		payload := map[string]interface{}{
			"user_id":    "12345",
			"event_type": "login",
			"timestamp":  time.Now().Format(time.RFC3339),
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshalling payload: %v\n", err)
			return
		}

		// Perform HTTP POST request to the API endpoint
		res, err := http.Post("http://localhost:8080/api/v1/user-event", "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		defer res.Body.Close()

		// Check if the request was successful
		if res.StatusCode == http.StatusOK {
			mutex.Lock()
			count++
			mutex.Unlock()
		}
	}

	// Main loop to send requests
	for {
		select {
		case <-stopChan:
			wg.Wait() // Wait for all goroutines to finish
			fmt.Printf("Total successful requests: %d\n", count)
			return
		default:
			wg.Add(1)
			// Block until there's a spot in the concurrency channel
			concurrencyChan <- struct{}{}
			go func() {
				defer func() { <-concurrencyChan }() // Release a spot in the concurrency channel
				request()
			}()
		}
	}
}
