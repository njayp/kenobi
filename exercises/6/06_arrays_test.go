package exercises

import (
	"reflect"
	"testing"
)

// Exercise: Arrays and Slices
// Task: Work with arrays and slices

// CreateArray: Create and return a fixed-size array of 5 integers with values [1, 2, 3, 4, 5]
func CreateArray() [5]int {
	// TODO: Create array with values [1, 2, 3, 4, 5]
	var arr [5]int
	return arr
}

// CreateSlice: Create and return a slice with values [10, 20, 30]
func CreateSlice() []int {
	// TODO: Create slice with values [10, 20, 30]
	var slice []int
	return slice
}

// SliceOperations: Perform various slice operations
func SliceOperations() ([]int, []int, int, int) {
	// TODO: Create slice [1, 2, 3, 4, 5, 6]
	s := []int{}
	
	// TODO: Create subslice from index 2 to 4 (exclusive) -> [3, 4]
	subSlice := []int{}
	
	// TODO: Append 7, 8 to original slice
	// TODO: Return: modified original slice, subslice, length, capacity
	return s, subSlice, 0, 0
}

// MakeSlice: Use make to create slice
func MakeSlice() []string {
	// TODO: Use make to create slice of strings with length 3, capacity 5
	// TODO: Set values to ["a", "b", "c"]
	return nil
}

func TestCreateArray(t *testing.T) {
	result := CreateArray()
	expected := [5]int{1, 2, 3, 4, 5}
	
	if result != expected {
		t.Errorf("CreateArray() = %v, want %v", result, expected)
	}
}

func TestCreateSlice(t *testing.T) {
	result := CreateSlice()
	expected := []int{10, 20, 30}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CreateSlice() = %v, want %v", result, expected)
	}
}

func TestSliceOperations(t *testing.T) {
	slice, subSlice, length, capacity := SliceOperations()
	
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedSubSlice := []int{3, 4}
	expectedLength := 8
	expectedCapacity := 12 // This might vary depending on Go version
	
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Errorf("Slice = %v, want %v", slice, expectedSlice)
	}
	if !reflect.DeepEqual(subSlice, expectedSubSlice) {
		t.Errorf("SubSlice = %v, want %v", subSlice, expectedSubSlice)
	}
	if length != expectedLength {
		t.Errorf("Length = %d, want %d", length, expectedLength)
	}
	if capacity < expectedLength {
		t.Errorf("Capacity = %d, should be at least %d", capacity, expectedLength)
	}
}

func TestMakeSlice(t *testing.T) {
	result := MakeSlice()
	expected := []string{"a", "b", "c"}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MakeSlice() = %v, want %v", result, expected)
	}
	if len(result) != 3 {
		t.Errorf("Length = %d, want 3", len(result))
	}
	if cap(result) != 5 {
		t.Errorf("Capacity = %d, want 5", cap(result))
	}
}
