package exercises

import (
	"testing"
)

// Exercise: Functions
// Task: Write functions with different signatures and behaviors

// SimpleFunction: A basic function that adds two numbers
func SimpleFunction(a, b int) int {
	// TODO: Return sum of a and b
	return 0
}

// MultipleReturnValues: Return multiple values
func MultipleReturnValues(a, b int) (int, int) {
	// TODO: Return sum and product of a and b
	return 0, 0
}

// NamedReturnValues: Use named return values
func NamedReturnValues(a, b int) (sum, product int) {
	// TODO: Set sum and product, then return (no explicit return values needed)
	return
}

// VariadicFunction: Accept variable number of arguments
func VariadicFunction(nums ...int) int {
	// TODO: Return sum of all numbers passed as arguments
	return 0
}

// HigherOrderFunction: Take a function as parameter
func HigherOrderFunction(fn func(int) int, value int) int {
	// TODO: Apply the function fn to value and return result
	return 0
}

// ReturnFunction: Return a function
func ReturnFunction(multiplier int) func(int) int {
	// TODO: Return a function that multiplies its input by multiplier
	return nil
}

func TestSimpleFunction(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{10, 5, 15},
		{-1, 1, 0},
	}
	
	for _, test := range tests {
		result := SimpleFunction(test.a, test.b)
		if result != test.expected {
			t.Errorf("SimpleFunction(%d, %d) = %d, want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMultipleReturnValues(t *testing.T) {
	sum, product := MultipleReturnValues(3, 4)
	
	if sum != 7 {
		t.Errorf("Sum = %d, want 7", sum)
	}
	if product != 12 {
		t.Errorf("Product = %d, want 12", product)
	}
}

func TestNamedReturnValues(t *testing.T) {
	sum, product := NamedReturnValues(5, 6)
	
	if sum != 11 {
		t.Errorf("Sum = %d, want 11", sum)
	}
	if product != 30 {
		t.Errorf("Product = %d, want 30", product)
	}
}

func TestVariadicFunction(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{10, 20}, 30},
		{[]int{5}, 5},
		{[]int{}, 0},
	}
	
	for _, test := range tests {
		result := VariadicFunction(test.nums...)
		if result != test.expected {
			t.Errorf("VariadicFunction(%v...) = %d, want %d", test.nums, result, test.expected)
		}
	}
}

func TestHigherOrderFunction(t *testing.T) {
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }
	
	if result := HigherOrderFunction(double, 5); result != 10 {
		t.Errorf("HigherOrderFunction(double, 5) = %d, want 10", result)
	}
	
	if result := HigherOrderFunction(square, 4); result != 16 {
		t.Errorf("HigherOrderFunction(square, 4) = %d, want 16", result)
	}
}

func TestReturnFunction(t *testing.T) {
	triple := ReturnFunction(3)
	if triple == nil {
		t.Fatal("ReturnFunction(3) returned nil")
	}
	
	if result := triple(4); result != 12 {
		t.Errorf("triple(4) = %d, want 12", result)
	}
	
	double := ReturnFunction(2)
	if result := double(7); result != 14 {
		t.Errorf("double(7) = %d, want 14", result)
	}
}
