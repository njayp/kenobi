package exercises

import (
	"reflect"
	"testing"
)

// Exercise: Pointers
// Task: Work with pointers and understand their behavior

// BasicPointer: Demonstrate basic pointer operations
func BasicPointer() (int, int, *int) {
	// TODO: Create an int variable with value 42
	i := 0
	
	// TODO: Create a pointer to i
	var p *int
	
	// TODO: Get value through pointer (dereferencing)
	val := 0
	
	// TODO: Return: original value, dereferenced value, and the pointer
	return i, val, p
}

// ModifyThroughPointer: Modify a value through its pointer
func ModifyThroughPointer(x *int) {
	// TODO: Set the value pointed to by x to 100
}

// SwapValues: Swap two values using pointers
func SwapValues(a, b *int) {
	// TODO: Swap the values that a and b point to
}

// NilPointer: Work with nil pointers safely
func NilPointer(p *int) int {
	// TODO: Check if pointer is nil, return 0 if nil, otherwise return the value
	return 0
}

func TestBasicPointer(t *testing.T) {
	original, dereferenced, pointer := BasicPointer()
	
	if original != 42 {
		t.Errorf("Original value = %d, want 42", original)
	}
	if dereferenced != 42 {
		t.Errorf("Dereferenced value = %d, want 42", dereferenced)
	}
	if pointer == nil {
		t.Error("Pointer is nil")
	}
	if *pointer != 42 {
		t.Errorf("Pointer points to %d, want 42", *pointer)
	}
}

func TestModifyThroughPointer(t *testing.T) {
	x := 50
	ModifyThroughPointer(&x)
	
	if x != 100 {
		t.Errorf("After ModifyThroughPointer, x = %d, want 100", x)
	}
}

func TestSwapValues(t *testing.T) {
	a, b := 10, 20
	SwapValues(&a, &b)
	
	if a != 20 {
		t.Errorf("After swap, a = %d, want 20", a)
	}
	if b != 10 {
		t.Errorf("After swap, b = %d, want 10", b)
	}
}

func TestNilPointer(t *testing.T) {
	// Test with nil pointer
	result := NilPointer(nil)
	if result != 0 {
		t.Errorf("NilPointer(nil) = %d, want 0", result)
	}
	
	// Test with valid pointer
	x := 42
	result = NilPointer(&x)
	if result != 42 {
		t.Errorf("NilPointer(&x) = %d, want 42", result)
	}
}
