package exercises

import (
	"testing"
)

// Exercise: For Loops
// Task: Implement different types of for loops

// BasicForLoop: Use a for loop to sum numbers 1 to n
func BasicForLoop(n int) int {
	// TODO: Use a for loop to calculate sum of 1 + 2 + ... + n
	sum := 0
	return sum
}

// WhileStyleLoop: Use a for loop as a while loop to count until condition is met
func WhileStyleLoop() int {
	// TODO: Use for loop as while loop to count from 1 until i > 5, return final i
	i := 1
	return i
}

// InfiniteBreakLoop: Use an infinite for loop with break to sum until sum > target
func InfiniteBreakLoop(target int) int {
	// TODO: Use infinite for loop, break when sum > target, return sum
	sum := 0
	i := 1
	return sum
}

func TestBasicForLoop(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{3, 6},   // 1+2+3
		{5, 15},  // 1+2+3+4+5
		{10, 55}, // 1+2+...+10
	}
	
	for _, test := range tests {
		result := BasicForLoop(test.n)
		if result != test.expected {
			t.Errorf("BasicForLoop(%d) = %d, want %d", test.n, result, test.expected)
		}
	}
}

func TestWhileStyleLoop(t *testing.T) {
	result := WhileStyleLoop()
	expected := 6 // Should stop when i becomes 6 (since 6 > 5)
	if result != expected {
		t.Errorf("WhileStyleLoop() = %d, want %d", result, expected)
	}
}

func TestInfiniteBreakLoop(t *testing.T) {
	result := InfiniteBreakLoop(10)
	expected := 15 // 1+2+3+4+5 = 15, first sum > 10
	if result != expected {
		t.Errorf("InfiniteBreakLoop(10) = %d, want %d", result, expected)
	}
}
