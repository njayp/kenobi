package exercises

import (
	"sort"
	"sync"
	"testing"
)

// Exercise: Mutexes
// Task: Use mutexes for thread-safe operations

// SafeCounter: Thread-safe counter using mutex
type SafeCounter struct {
	// TODO: Add mutex and counter fields
}

// NewSafeCounter: Constructor for SafeCounter
func NewSafeCounter() *SafeCounter {
	// TODO: Return pointer to new SafeCounter
	return nil
}

// Increment: Thread-safe increment
func (c *SafeCounter) Increment() {
	// TODO: Use mutex to safely increment counter
}

// Decrement: Thread-safe decrement
func (c *SafeCounter) Decrement() {
	// TODO: Use mutex to safely decrement counter
}

// Value: Thread-safe getter
func (c *SafeCounter) Value() int {
	// TODO: Use mutex to safely read counter value
	return 0
}

// SafeMap: Thread-safe map using RWMutex
type SafeMap struct {
	// TODO: Add RWMutex and map[string]int fields
}

// NewSafeMap: Constructor for SafeMap
func NewSafeMap() *SafeMap {
	// TODO: Return pointer to new SafeMap with initialized map
	return nil
}

// Set: Thread-safe setter
func (sm *SafeMap) Set(key string, value int) {
	// TODO: Use write lock to safely set value
}

// Get: Thread-safe getter
func (sm *SafeMap) Get(key string) (int, bool) {
	// TODO: Use read lock to safely get value
	return 0, false
}

// Keys: Return all keys (thread-safe)
func (sm *SafeMap) Keys() []string {
	// TODO: Use read lock to safely get all keys
	return nil
}

// ConcurrentAccess: Demonstrate concurrent access to shared resource
func ConcurrentAccess() int {
	// TODO: Create SafeCounter
	// TODO: Launch 10 goroutines, each increments counter 100 times
	// TODO: Use WaitGroup to wait for all goroutines
	// TODO: Return final counter value (should be 1000)
	return 0
}

// RaceConditionDemo: Show what happens without mutex (for educational purposes)
func RaceConditionDemo() int {
	// TODO: Create regular int counter (no mutex)
	// TODO: Launch 10 goroutines, each increments counter 100 times
	// TODO: Use WaitGroup to wait for all goroutines
	// TODO: Return final counter value (will likely be less than 1000 due to race condition)
	counter := 0
	var wg sync.WaitGroup
	
	// TODO: Implement race condition demonstration
	
	return counter
}

func TestSafeCounter(t *testing.T) {
	counter := NewSafeCounter()
	
	if counter == nil {
		t.Fatal("NewSafeCounter() returned nil")
	}
	
	// Test initial value
	if counter.Value() != 0 {
		t.Errorf("Initial value = %d, want 0", counter.Value())
	}
	
	// Test increment
	counter.Increment()
	if counter.Value() != 1 {
		t.Errorf("After increment, value = %d, want 1", counter.Value())
	}
	
	// Test decrement
	counter.Decrement()
	if counter.Value() != 0 {
		t.Errorf("After decrement, value = %d, want 0", counter.Value())
	}
}

func TestSafeMap(t *testing.T) {
	sm := NewSafeMap()
	
	if sm == nil {
		t.Fatal("NewSafeMap() returned nil")
	}
	
	// Test set and get
	sm.Set("key1", 10)
	value, ok := sm.Get("key1")
	if !ok {
		t.Error("Get(\"key1\") should return true")
	}
	if value != 10 {
		t.Errorf("Get(\"key1\") = %d, want 10", value)
	}
	
	// Test non-existent key
	_, ok = sm.Get("nonexistent")
	if ok {
		t.Error("Get(\"nonexistent\") should return false")
	}
	
	// Test multiple keys
	sm.Set("key2", 20)
	sm.Set("key3", 30)
	
	keys := sm.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys() returned %d keys, want 3", len(keys))
	}
	
	// Sort keys for consistent testing
	sort.Strings(keys)
	expected := []string{"key1", "key2", "key3"}
	for i, expectedKey := range expected {
		if keys[i] != expectedKey {
			t.Errorf("Keys()[%d] = %q, want %q", i, keys[i], expectedKey)
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	result := ConcurrentAccess()
	expected := 1000 // 10 goroutines * 100 increments each
	
	if result != expected {
		t.Errorf("ConcurrentAccess() = %d, want %d", result, expected)
	}
}

func TestRaceConditionDemo(t *testing.T) {
	// Run multiple times to increase chance of seeing race condition
	var results []int
	for i := 0; i < 5; i++ {
		result := RaceConditionDemo()
		results = append(results, result)
	}
	
	// At least one result should be less than 1000 due to race condition
	// (Though this is not guaranteed in all environments)
	allEqual := true
	for _, result := range results {
		if result != 1000 {
			allEqual = false
			break
		}
	}
	
	// If all results are 1000, that's actually fine for this test
	// The important thing is that students understand the concept
	t.Logf("Race condition demo results: %v", results)
	if allEqual {
		t.Log("No race condition detected in this run, but that's normal on some systems")
	}
}

// Benchmark to show performance difference between mutex and RWMutex
func BenchmarkSafeMapWrite(b *testing.B) {
	sm := NewSafeMap()
	if sm == nil {
		b.Fatal("NewSafeMap() returned nil")
	}
	
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Set("key", i)
			i++
		}
	})
}

func BenchmarkSafeMapRead(b *testing.B) {
	sm := NewSafeMap()
	if sm == nil {
		b.Fatal("NewSafeMap() returned nil")
	}
	
	sm.Set("key", 42)
	
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sm.Get("key")
		}
	})
}
