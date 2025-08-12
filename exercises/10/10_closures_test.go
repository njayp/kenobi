package exercises

import (
	"testing"
)

// Exercise: Closures
// Task: Understand and implement closures

// Counter: Return a closure that increments and returns a counter
func Counter1() func() int {
	// TODO: Create a counter variable and return function that increments and returns it
	return nil
}

// Adder: Return a closure that adds a specific value to its input
func Adder(addend int) func(int) int {
	// TODO: Return a function that adds addend to its parameter
	return nil
}

// Accumulator: Return a closure that accumulates values
func Accumulator() func(int) int {
	// TODO: Return a function that adds each input to a running total and returns the total
	return nil
}

func TestCounter(t *testing.T) {
	counter1 := Counter1()
	counter2 := Counter1()

	if counter1 == nil {
		t.Fatal("Counter() returned nil")
	}

	// Test first counter
	if result := counter1(); result != 1 {
		t.Errorf("counter1() first call = %d, want 1", result)
	}
	if result := counter1(); result != 2 {
		t.Errorf("counter1() second call = %d, want 2", result)
	}
	if result := counter1(); result != 3 {
		t.Errorf("counter1() third call = %d, want 3", result)
	}

	// Test second counter (should be independent)
	if result := counter2(); result != 1 {
		t.Errorf("counter2() first call = %d, want 1", result)
	}
}

func TestAdder(t *testing.T) {
	add5 := Adder(5)
	add10 := Adder(10)

	if add5 == nil || add10 == nil {
		t.Fatal("Adder() returned nil")
	}

	if result := add5(3); result != 8 {
		t.Errorf("add5(3) = %d, want 8", result)
	}
	if result := add5(7); result != 12 {
		t.Errorf("add5(7) = %d, want 12", result)
	}

	if result := add10(3); result != 13 {
		t.Errorf("add10(3) = %d, want 13", result)
	}
}

func TestAccumulator(t *testing.T) {
	acc := Accumulator()

	if acc == nil {
		t.Fatal("Accumulator() returned nil")
	}

	if result := acc(5); result != 5 {
		t.Errorf("acc(5) = %d, want 5", result)
	}
	if result := acc(3); result != 8 {
		t.Errorf("acc(3) = %d, want 8", result)
	}
	if result := acc(2); result != 10 {
		t.Errorf("acc(2) = %d, want 10", result)
	}
}
