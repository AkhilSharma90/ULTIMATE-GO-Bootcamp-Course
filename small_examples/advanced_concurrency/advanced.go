package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

// Producer: Sends logs into a buffered channel
func logProducer(logs []LogEntry, logChan chan<- LogEntry) {
	for _, log := range logs {
		logChan <- log
		time.Sleep(100 * time.Millisecond) // Simulate delay in log generation
	}
	close(logChan)
}

// Filter: Filters logs based on level and sends them to the filtered channel
func logFilter(logChan <-chan LogEntry, filteredChan chan<- LogEntry, level string, wg *sync.WaitGroup) {
	defer wg.Done()
	for log := range logChan {
		if log.Level == level {
			filteredChan <- log
		}
	}
}

// Aggregator: Aggregates statistics from filtered logs and sends them via stats channel
func logAggregator(filteredChan <-chan LogEntry, statsChan chan<- map[string]int, wg sync.WaitGroup, mu sync.Mutex) {
	defer wg.Done()
	stats := make(map[string]int)
	for log := range filteredChan {
		mu.Lock() // Protect shared resource
		stats[log.Level]++
		mu.Unlock()
	}
	// Send aggregated stats
	statsChan <- stats
	close(statsChan)
}

// Consumer: Reads from stats channel and uses select for dynamic handling
func statsConsumer(statsChan <-chan map[string]int, done chan<- bool) {
	for {
		select {
		case stats, ok := <-statsChan:
			if !ok {
				fmt.Println("Stats processing complete")
				done <- true
				return
			}
			for level, count := range stats {
				fmt.Printf("Level: %s, Count: %d\n", strings.ToUpper(level), count)
			}
		case <-time.After(3 * time.Second): // Timeout handling
			fmt.Println("Timeout: No stats received in 3 seconds")
			done <- true
			return
		}
	}
}
func main() {
	// Simulated log entries
	logs := []LogEntry{
		{"2023-01-01T10:00:00Z", "INFO", "Service started"},
		{"2023-01-01T10:05:00Z", "ERROR", "Database connection failed"},
		{"2023-01-01T10:10:00Z", "INFO", "Service running"},
		{"2023-01-01T10:15:00Z", "ERROR", "Disk space low"},
		{"2023-01-01T10:20:00Z", "INFO", "Service stopped"},
	}
	logChan := make(chan LogEntry, 5)         // Buffered channel for logs
	filteredChan := make(chan LogEntry)       // Unbuffered channel for filtered logs
	statsChan := make(chan map[string]int, 1) // Buffered channel for stats
	done := make(chan bool)                   // Signal channel for completion
	var wg sync.WaitGroup
	var mu sync.Mutex
	// Start the log producer
	go logProducer(logs, logChan)
	// Start filter workers
	numFilters := 2
	wg.Add(numFilters)
	for i := 0; i < numFilters; i++ {
		go logFilter(logChan, filteredChan, "ERROR", &wg)
	}
	// Close filtered channel once all filter workers are done
	go func() {
		wg.Wait()
		close(filteredChan)
	}()
	// Start the aggregator
	wg.Add(1)
	go logAggregator(filteredChan, statsChan, &wg, &mu)
	// Start the stats consumer
	go statsConsumer(statsChan, done)
	// Wait for stats processing to complete
	<-done
	fmt.Println("All log processing tasks completed")
}
