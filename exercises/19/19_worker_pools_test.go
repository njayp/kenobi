package exercises

import (
	"sync"
	"testing"
	"time"
)

// Exercise: Worker Pools
// Task: Implement worker pool pattern with goroutines and channels

// Job represents work to be done
type Job struct {
	ID   int
	Data int
}

// Result represents the result of processing a job
type Result struct {
	JobID int
	Value int
}

// WorkerPool: Implement a worker pool that processes jobs
func WorkerPool(numWorkers int, jobs []Job) []Result {
	// TODO: Create channels for jobs and results
	// TODO: Launch numWorkers goroutines (workers)
	// TODO: Each worker should:
	//   - Receive jobs from jobs channel
	//   - Process job (square the Data field)
	//   - Send result to results channel
	// TODO: Send all jobs to jobs channel and close it
	// TODO: Collect all results and return
	return nil
}

// worker: Helper function for processing jobs
func worker(id int, jobs <-chan Job, results chan<- Result) {
	// TODO: Implement worker that processes jobs from channel
	// TODO: For each job, square the Data and send Result
}

// Producer-Consumer pattern
func ProducerConsumer() []int {
	// TODO: Create channel for communication
	// TODO: Launch producer goroutine that sends numbers 1-10
	// TODO: Launch consumer goroutine that receives and processes numbers
	// TODO: Return slice of processed numbers (each number doubled)
	return nil
}

// RateLimiter: Implement rate limiting using channels
func RateLimiter(requests int) time.Duration {
	// TODO: Create ticker that ticks every 100ms (rate limit)
	// TODO: Process 'requests' number of requests
	// TODO: Each request should wait for ticker before proceeding
	// TODO: Return total time taken
	start := time.Now()
	
	// TODO: Implement rate limiting logic
	
	return time.Since(start)
}

// FanOut-FanIn pattern
func FanOutFanIn(input []int) []int {
	// TODO: Create input channel and send all input values
	// TODO: Create multiple worker goroutines (fan-out)
	// TODO: Each worker processes input (multiply by 2)
	// TODO: Merge results from all workers (fan-in)
	// TODO: Return collected results
	return nil
}

func TestWorkerPool(t *testing.T) {
	jobs := []Job{
		{ID: 1, Data: 2},
		{ID: 2, Data: 3},
		{ID: 3, Data: 4},
		{ID: 4, Data: 5},
	}
	
	results := WorkerPool(2, jobs)
	
	if len(results) != len(jobs) {
		t.Errorf("Expected %d results, got %d", len(jobs), len(results))
		return
	}
	
	// Create map to check results
	resultMap := make(map[int]int)
	for _, result := range results {
		resultMap[result.JobID] = result.Value
	}
	
	// Check that each job was processed correctly (squared)
	expected := map[int]int{
		1: 4,  // 2²
		2: 9,  // 3²
		3: 16, // 4²
		4: 25, // 5²
	}
	
	for jobID, expectedValue := range expected {
		if resultMap[jobID] != expectedValue {
			t.Errorf("Job %d: got %d, want %d", jobID, resultMap[jobID], expectedValue)
		}
	}
}

func TestProducerConsumer(t *testing.T) {
	results := ProducerConsumer()
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} // 1*2, 2*2, ..., 10*2
	
	if len(results) != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), len(results))
		return
	}
	
	for i, expectedValue := range expected {
		if results[i] != expectedValue {
			t.Errorf("Index %d: got %d, want %d", i, results[i], expectedValue)
		}
	}
}

func TestRateLimiter(t *testing.T) {
	// Test with 3 requests, should take at least 200ms (3 requests, 100ms each after first)
	duration := RateLimiter(3)
	
	expectedMin := 200 * time.Millisecond
	if duration < expectedMin {
		t.Errorf("Duration = %v, expected at least %v", duration, expectedMin)
	}
	
	// Should not take too much longer (allow some overhead)
	expectedMax := 350 * time.Millisecond
	if duration > expectedMax {
		t.Errorf("Duration = %v, expected less than %v", duration, expectedMax)
	}
}

func TestFanOutFanIn(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	results := FanOutFanIn(input)
	
	if len(results) != len(input) {
		t.Errorf("Expected %d results, got %d", len(input), len(results))
		return
	}
	
	// Results might be in different order due to concurrent processing
	// So we'll check that all expected values are present
	expected := []int{2, 4, 6, 8, 10} // input values * 2
	
	resultMap := make(map[int]bool)
	for _, v := range results {
		resultMap[v] = true
	}
	
	for _, expectedValue := range expected {
		if !resultMap[expectedValue] {
			t.Errorf("Missing expected value %d in results %v", expectedValue, results)
		}
	}
}

// Bonus: Pipeline pattern
func Pipeline() int {
	// TODO: Implement pipeline with multiple stages
	// Stage 1: Generate numbers 1-5
	// Stage 2: Square each number  
	// Stage 3: Sum all squared numbers
	// Return final sum
	return 0
}

func TestPipeline(t *testing.T) {
	result := Pipeline()
	expected := 55 // 1² + 2² + 3² + 4² + 5² = 1 + 4 + 9 + 16 + 25 = 55
	
	if result != expected {
		t.Errorf("Pipeline() = %d, want %d", result, expected)
	}
}
