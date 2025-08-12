package exercises

import (
	"testing"
)

// Exercise: If/Else Statements
// Task: Implement conditional logic

// BasicIfElse: Return "positive", "negative", or "zero" based on number
func BasicIfElse(n int) string {
	// TODO: Return "positive" if n > 0, "negative" if n < 0, "zero" if n == 0
	return ""
}

// IfWithoutElse: Return true if number is even, false otherwise
func IfWithoutElse(n int) bool {
	// TODO: Return true if n is even (n % 2 == 0)
	return false
}

// IfElseIfElse: Categorize a number as "small", "medium", or "large"
func IfElseIfElse(n int) string {
	// TODO: Return "small" if n < 10, "medium" if n < 100, "large" otherwise
	return ""
}

func TestBasicIfElse(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{5, "positive"},
		{-3, "negative"},
		{0, "zero"},
		{100, "positive"},
		{-1, "negative"},
	}
	
	for _, test := range tests {
		result := BasicIfElse(test.n)
		if result != test.expected {
			t.Errorf("BasicIfElse(%d) = %q, want %q", test.n, result, test.expected)
		}
	}
}

func TestIfWithoutElse(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{7, false},
	}
	
	for _, test := range tests {
		result := IfWithoutElse(test.n)
		if result != test.expected {
			t.Errorf("IfWithoutElse(%d) = %t, want %t", test.n, result, test.expected)
		}
	}
}

func TestIfElseIfElse(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{5, "small"},
		{50, "medium"},
		{500, "large"},
		{9, "small"},
		{99, "medium"},
		{100, "large"},
	}
	
	for _, test := range tests {
		result := IfElseIfElse(test.n)
		if result != test.expected {
			t.Errorf("IfElseIfElse(%d) = %q, want %q", test.n, result, test.expected)
		}
	}
}
