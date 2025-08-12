package exercises

import (
	"testing"
)

// Exercise: Range
// Task: Use range to iterate over different data structures

// RangeSlice: Use range to sum all numbers in a slice
func RangeSlice(numbers []int) int {
	// TODO: Use range to iterate and sum all numbers
	sum := 0
	return sum
}

// RangeSliceIndexValue: Use range to find index of target value
func RangeSliceIndexValue(slice []string, target string) int {
	// TODO: Use range to find index of target string, return -1 if not found
	return -1
}

// RangeMap: Use range to collect all keys from a map
func RangeMap(m map[string]int) []string {
	// TODO: Use range to collect all keys into a slice
	var keys []string
	return keys
}

// RangeString: Use range to count vowels in a string
func RangeString(s string) int {
	// TODO: Use range to iterate over string and count vowels (a, e, i, o, u)
	count := 0
	return count
}

func TestRangeSlice(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{10, 20, 30}, 60},
		{[]int{}, 0},
		{[]int{-1, 2, -3}, -2},
	}
	
	for _, test := range tests {
		result := RangeSlice(test.numbers)
		if result != test.expected {
			t.Errorf("RangeSlice(%v) = %d, want %d", test.numbers, result, test.expected)
		}
	}
}

func TestRangeSliceIndexValue(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date"}
	
	tests := []struct {
		target   string
		expected int
	}{
		{"apple", 0},
		{"cherry", 2},
		{"date", 3},
		{"grape", -1},
	}
	
	for _, test := range tests {
		result := RangeSliceIndexValue(slice, test.target)
		if result != test.expected {
			t.Errorf("RangeSliceIndexValue(%v, %q) = %d, want %d", slice, test.target, result, test.expected)
		}
	}
}

func TestRangeMap(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	
	result := RangeMap(m)
	
	if len(result) != 3 {
		t.Errorf("RangeMap() returned %d keys, want 3", len(result))
	}
	
	// Check that all expected keys are present
	keyMap := make(map[string]bool)
	for _, key := range result {
		keyMap[key] = true
	}
	
	for expectedKey := range m {
		if !keyMap[expectedKey] {
			t.Errorf("RangeMap() missing key %q", expectedKey)
		}
	}
}

func TestRangeString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 2},     // e, o
		{"programming", 3}, // o, a, i
		{"xyz", 0},        // no vowels
		{"aeiou", 5},      // all vowels
		{"", 0},           // empty string
	}
	
	for _, test := range tests {
		result := RangeString(test.input)
		if result != test.expected {
			t.Errorf("RangeString(%q) = %d, want %d", test.input, result, test.expected)
		}
	}
}
