package exercises

import (
	"testing"
)

// Exercise: Methods
// Task: Define methods for custom types

// Define a custom type for demonstrating methods
type Counter struct {
	// TODO: Add a count field of type int
}

// NewCounter: Constructor for Counter
func NewCounter() *Counter {
	// TODO: Return pointer to new Counter with count initialized to 0
	return nil
}

// Increment: Method to increment counter (pointer receiver)
func (c *Counter) Increment() {
	// TODO: Increment the count by 1
}

// Decrement: Method to decrement counter (pointer receiver)
func (c *Counter) Decrement() {
	// TODO: Decrement the count by 1
}

// Value: Method to get current value (value receiver)
func (c Counter) Value() int {
	// TODO: Return the current count
	return 0
}

// Reset: Method to reset counter to 0 (pointer receiver)
func (c *Counter) Reset() {
	// TODO: Set count to 0
}

// String type for string methods
type String string

// Length: Method on String type to get length
func (s String) Length() int {
	// TODO: Return length of string
	return 0
}

// Upper: Method to convert to uppercase
func (s String) Upper() string {
	// TODO: Convert string to uppercase and return
	return ""
}

// Contains: Method to check if string contains substring
func (s String) Contains(substr string) bool {
	// TODO: Check if s contains substr
	return false
}

func TestCounterMethods(t *testing.T) {
	counter := NewCounter()
	
	if counter == nil {
		t.Fatal("NewCounter() returned nil")
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
	
	// Test multiple increments
	counter.Increment()
	counter.Increment()
	if counter.Value() != 3 {
		t.Errorf("After 3 increments, value = %d, want 3", counter.Value())
	}
	
	// Test decrement
	counter.Decrement()
	if counter.Value() != 2 {
		t.Errorf("After decrement, value = %d, want 2", counter.Value())
	}
	
	// Test reset
	counter.Reset()
	if counter.Value() != 0 {
		t.Errorf("After reset, value = %d, want 0", counter.Value())
	}
}

func TestStringMethods(t *testing.T) {
	s := String("Hello World")
	
	// Test Length
	if length := s.Length(); length != 11 {
		t.Errorf("Length() = %d, want 11", length)
	}
	
	// Test Upper
	if upper := s.Upper(); upper != "HELLO WORLD" {
		t.Errorf("Upper() = %q, want %q", upper, "HELLO WORLD")
	}
	
	// Test Contains
	if !s.Contains("World") {
		t.Error("Contains(\"World\") = false, want true")
	}
	if s.Contains("xyz") {
		t.Error("Contains(\"xyz\") = true, want false")
	}
}

// Test method sets (value vs pointer receivers)
func TestMethodSets(t *testing.T) {
	// Test that pointer receivers work with both pointer and value
	counter := NewCounter()
	counter.Increment() // pointer receiver method
	
	// Test that value receivers work with both pointer and value
	value := counter.Value() // value receiver method
	if value != 1 {
		t.Errorf("Method call on pointer, Value() = %d, want 1", value)
	}
	
	// Test value receiver on value
	counterValue := *counter
	value = counterValue.Value()
	if value != 1 {
		t.Errorf("Method call on value, Value() = %d, want 1", value)
	}
}
