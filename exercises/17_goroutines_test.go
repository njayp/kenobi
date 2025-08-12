package exercises

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Exercise: Goroutines
// Task: Work with goroutines for concurrent programming

// SimpleGoroutine: Launch a goroutine that prints numbers
func SimpleGoroutine() {
	// TODO: Use go keyword to launch a goroutine that prints numbers 1-5
	// TODO: Use time.Sleep to allow goroutine to complete
}

// GoroutineWithChannel: Use goroutine with channel for communication
func GoroutineWithChannel() string {
	// TODO: Create a channel of strings
	// TODO: Launch goroutine that sends "Hello from goroutine" to channel
	// TODO: Receive and return the message from channel
	return ""
}

// MultipleGoroutines: Launch multiple goroutines
func MultipleGoroutines(n int) []int {
	// TODO: Create slice to store results
	// TODO: Create channel for receiving results
	// TODO: Launch n goroutines, each sends its number (i) to channel
	// TODO: Collect all results and return sorted slice
	return nil
}

// GoroutineWithWaitGroup: Use sync.WaitGroup to wait for goroutines
func GoroutineWithWaitGroup() int {
	// TODO: Create WaitGroup
	// TODO: Create shared counter (use mutex for thread safety)
	// TODO: Launch 5 goroutines that each increment counter 100 times
	// TODO: Wait for all goroutines to complete
	// TODO: Return final counter value (should be 500)
	return 0
}

func TestSimpleGoroutine(t *testing.T) {
	// This test just verifies the function doesn't panic
	// In a real scenario, you'd capture output or use channels
	SimpleGoroutine()
}

func TestGoroutineWithChannel(t *testing.T) {
	result := GoroutineWithChannel()
	expected := "Hello from goroutine"
	
	if result != expected {
		t.Errorf("GoroutineWithChannel() = %q, want %q", result, expected)
	}
}

func TestMultipleGoroutines(t *testing.T) {
	result := MultipleGoroutines(3)
	expected := []int{0, 1, 2}
	
	if len(result) != len(expected) {
		t.Errorf("Length = %d, want %d", len(result), len(expected))
		return
	}
	
	// Check if all expected numbers are present (order might vary)
	resultMap := make(map[int]bool)
	for _, v := range result {
		resultMap[v] = true
	}
	
	for _, expected := range expected {
		if !resultMap[expected] {
			t.Errorf("Missing value %d in result %v", expected, result)
		}
	}
}

func TestGoroutineWithWaitGroup(t *testing.T) {
	result := GoroutineWithWaitGroup()
	expected := 500
	
	if result != expected {
		t.Errorf("GoroutineWithWaitGroup() = %d, want %d", result, expected)
	}
}
