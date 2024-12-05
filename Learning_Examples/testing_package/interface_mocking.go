package main

import "testing"

// Dependency Interface
type Fetcher interface {
	Fetch() string
}

// Mock Implementation
type MockFetcher struct{}

func (m MockFetcher) Fetch() string {
	return "mock response"
}

// Function to Test
func GetData(fetcher Fetcher) string {
	return fetcher.Fetch()
}

func TestGetData(t *testing.T) {
	mock := MockFetcher{}
	result := GetData(mock)
	if result != "mock response" {
		t.Errorf("Expected 'mock response', got '%s'", result)
	}
}
