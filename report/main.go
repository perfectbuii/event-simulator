package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand/v2"
	"os"
)

// PerformanceData represents the structure of the performance test data
type PerformanceData struct {
	Tab             string
	IterationRate   float64
	RequestRate     float64
	RequestDuration float64
	RequestFailed   float64
	SentRate        float64
	Performance     []DataPoint
	VUS             []DataPoint
	TransferRate    []DataPoint
	RequestDur      []DataPoint
	IterationDur    []DataPoint
	PerformanceJS   string
	VUSJS           string
	TransferRateJS  string
	RequestDurJS    string
	IterationDurJS  string
}

// DataPoint represents a point in the chart
type DataPoint struct {
	X float64
	Y float64
}

// ChartData holds all tabs' performance data and the active tab
type ChartData struct {
	Tabs      []PerformanceData
	ActiveTab string
}

// Convert a slice of DataPoint to a JSON array for JavaScript
func dataPointsToJSArray(data []DataPoint) string {
	jsArray, err := json.Marshal(struct {
		Labels []float64 `json:"labels"`
		Values []float64 `json:"values"`
	}{
		Labels: extractLabels(data),
		Values: extractValues(data),
	})
	if err != nil {
		log.Fatalf("Failed to marshal data points: %s", err)
	}
	return string(jsArray)
}

// Extract labels from DataPoint
func extractLabels(data []DataPoint) []float64 {
	labels := make([]float64, len(data))
	for i, dp := range data {
		labels[i] = dp.X
	}
	return labels
}

// Extract values from DataPoint
func extractValues(data []DataPoint) []float64 {
	values := make([]float64, len(data))
	for i, dp := range data {
		values[i] = dp.Y
	}
	return values
}

func main() {
	// Init data
	data := ChartData{
		Tabs: []PerformanceData{
			{
				Tab:             "HTTP",
				IterationRate:   randomFloat(50, 200),
				RequestRate:     randomFloat(100, 300),
				RequestDuration: randomFloat(200, 400),
				RequestFailed:   randomFloat(1, 10),
				SentRate:        randomFloat(400, 700),
				Performance:     generateData(10),
				VUS:             generateData(10),
				TransferRate:    generateData(10),
				RequestDur:      generateData(10),
				IterationDur:    generateData(10),
			},
			{
				Tab:             "GRPC",
				IterationRate:   randomFloat(50, 200),
				RequestRate:     randomFloat(100, 300),
				RequestDuration: randomFloat(200, 400),
				RequestFailed:   randomFloat(1, 10),
				SentRate:        randomFloat(400, 700),
				Performance:     generateData(10),
				VUS:             generateData(10),
				TransferRate:    generateData(10),
				RequestDur:      generateData(10),
				IterationDur:    generateData(10),
			},
			{
				Tab:             "Kafka",
				IterationRate:   randomFloat(50, 200),
				RequestRate:     randomFloat(100, 300),
				RequestDuration: randomFloat(200, 400),
				RequestFailed:   randomFloat(1, 10),
				SentRate:        randomFloat(400, 700),
				Performance:     generateData(10),
				VUS:             generateData(10),
				TransferRate:    generateData(10),
				RequestDur:      generateData(10),
				IterationDur:    generateData(10),
			},
		},
		ActiveTab: "HTTP", // Set the default active tab here
	}

	// Preprocess data for JavaScript
	for i := range data.Tabs {
		tab := &data.Tabs[i]
		tab.PerformanceJS = dataPointsToJSArray(tab.Performance)
		tab.VUSJS = dataPointsToJSArray(tab.VUS)
		tab.TransferRateJS = dataPointsToJSArray(tab.TransferRate)
		tab.RequestDurJS = dataPointsToJSArray(tab.RequestDur)
		tab.IterationDurJS = dataPointsToJSArray(tab.IterationDur)
	}

	// Define the function map
	funcMap := template.FuncMap{}

	// Parse the template file with the function map
	tmpl := template.Must(template.New("chart.tmpl").Funcs(funcMap).ParseFiles("chart.tmpl"))

	// Create a file to write the output to
	file, err := os.Create("chart.html")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	// Execute the template with the data
	err = tmpl.Execute(file, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}

	log.Println("chart.html has been successfully created.")
}

// Function to generate dummy data
func generateData(n int) []DataPoint {
	data := make([]DataPoint, n)
	for i := 0; i < n; i++ {
		data[i] = DataPoint{
			X: float64(i),
			Y: randomFloat(0, 100), // Random Y value between 0 and 100
		}
	}
	return data
}

// Helper function to generate random float values
func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
