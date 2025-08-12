package exercises

import (
	"testing"
)

// Exercise: Recursion
// Task: Implement recursive functions

// Factorial: Calculate factorial of n using recursion
func Factorial(n int) int {
	// TODO: Implement factorial recursively
	// Base case: factorial(0) = 1, factorial(1) = 1
	// Recursive case: factorial(n) = n * factorial(n-1)
	return 0
}

// Fibonacci: Calculate nth Fibonacci number using recursion
func Fibonacci(n int) int {
	// TODO: Implement Fibonacci recursively
	// Base cases: fib(0) = 0, fib(1) = 1
	// Recursive case: fib(n) = fib(n-1) + fib(n-2)
	return 0
}

// SumDigits: Calculate sum of digits in a number using recursion
func SumDigits(n int) int {
	// TODO: Implement sum of digits recursively
	// Base case: if n < 10, return n
	// Recursive case: return (n % 10) + SumDigits(n / 10)
	return 0
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}
	
	for _, test := range tests {
		result := Factorial(test.n)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d, want %d", test.n, result, test.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}
	
	for _, test := range tests {
		result := Fibonacci(test.n)
		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d, want %d", test.n, result, test.expected)
		}
	}
}

func TestSumDigits(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{5, 5},
		{12, 3},    // 1 + 2
		{123, 6},   // 1 + 2 + 3
		{999, 27},  // 9 + 9 + 9
		{1234, 10}, // 1 + 2 + 3 + 4
	}
	
	for _, test := range tests {
		result := SumDigits(test.n)
		if result != test.expected {
			t.Errorf("SumDigits(%d) = %d, want %d", test.n, result, test.expected)
		}
	}
}
