package exercises

import (
	"testing"
)

// Exercise: Interfaces
// Task: Define and implement interfaces

// Shape interface defines methods that geometric shapes should implement
type Shape interface {
	// TODO: Add Area() method that returns float64
	// TODO: Add Perimeter() method that returns float64
}

// Circle struct
type Circle struct {
	// TODO: Add Radius field of type float64
}

// Area method for Circle (implement Shape interface)
func (c Circle) Area() float64 {
	// TODO: Return π * r² (use 3.14159 for π)
	return 0
}

// Perimeter method for Circle (implement Shape interface)
func (c Circle) Perimeter() float64 {
	// TODO: Return 2 * π * r (use 3.14159 for π)
	return 0
}

// Square struct
type Square struct {
	// TODO: Add Side field of type float64
}

// Area method for Square (implement Shape interface)
func (s Square) Area() float64 {
	// TODO: Return side²
	return 0
}

// Perimeter method for Square (implement Shape interface)
func (s Square) Perimeter() float64 {
	// TODO: Return 4 * side
	return 0
}

// CalculateTotalArea: Function that works with any Shape
func CalculateTotalArea(shapes []Shape) float64 {
	// TODO: Calculate and return sum of areas of all shapes
	return 0
}

// Stringer interface (similar to fmt.Stringer)
type Stringer interface {
	String() string
}

// Point struct that implements Stringer
type Point struct {
	X, Y int
}

// String method for Point (implement Stringer interface)
func (p Point) String() string {
	// TODO: Return string in format "(x, y)"
	return ""
}

// PrintString: Function that works with any Stringer
func PrintString(s Stringer) string {
	// TODO: Return s.String()
	return ""
}

func TestShapeInterface(t *testing.T) {
	circle := Circle{Radius: 5}
	square := Square{Side: 4}
	
	// Test Circle methods
	if area := circle.Area(); area != 78.53975 { // π * 5²
		t.Errorf("Circle.Area() = %f, want 78.53975", area)
	}
	if perimeter := circle.Perimeter(); perimeter != 31.4159 { // 2 * π * 5
		t.Errorf("Circle.Perimeter() = %f, want 31.4159", perimeter)
	}
	
	// Test Square methods
	if area := square.Area(); area != 16 { // 4²
		t.Errorf("Square.Area() = %f, want 16", area)
	}
	if perimeter := square.Perimeter(); perimeter != 16 { // 4 * 4
		t.Errorf("Square.Perimeter() = %f, want 16", perimeter)
	}
}

func TestPolymorphism(t *testing.T) {
	// Test that Circle and Square can be used as Shape interface
	var shape1 Shape = Circle{Radius: 2}
	var shape2 Shape = Square{Side: 3}
	
	if area := shape1.Area(); area != 12.56636 { // π * 2²
		t.Errorf("Circle as Shape, Area() = %f, want 12.56636", area)
	}
	if area := shape2.Area(); area != 9 { // 3²
		t.Errorf("Square as Shape, Area() = %f, want 9", area)
	}
}

func TestCalculateTotalArea(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 1},   // Area: π = 3.14159
		Square{Side: 2},     // Area: 4
		Circle{Radius: 2},   // Area: 4π = 12.56636
	}
	
	total := CalculateTotalArea(shapes)
	expected := 19.70795 // 3.14159 + 4 + 12.56636
	
	if total != expected {
		t.Errorf("CalculateTotalArea() = %f, want %f", total, expected)
	}
}

func TestStringerInterface(t *testing.T) {
	point := Point{X: 3, Y: 4}
	
	// Test String method
	if str := point.String(); str != "(3, 4)" {
		t.Errorf("Point.String() = %q, want %q", str, "(3, 4)")
	}
	
	// Test using interface
	var stringer Stringer = point
	if str := stringer.String(); str != "(3, 4)" {
		t.Errorf("Point as Stringer, String() = %q, want %q", str, "(3, 4)")
	}
	
	// Test PrintString function
	if str := PrintString(point); str != "(3, 4)" {
		t.Errorf("PrintString(point) = %q, want %q", str, "(3, 4)")
	}
}

// Test empty interface
func TestEmptyInterface(t *testing.T) {
	// TODO: Implement a function that accepts empty interface and returns type name
	// This is for students to learn about interface{} and type assertions
}

func GetTypeName(i interface{}) string {
	// TODO: Use type switch to return type name as string
	// Return "int", "string", "bool", or "unknown" for other types
	return ""
}

func TestGetTypeName(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{"hello", "string"},
		{true, "bool"},
		{3.14, "unknown"},
	}
	
	for _, test := range tests {
		result := GetTypeName(test.input)
		if result != test.expected {
			t.Errorf("GetTypeName(%v) = %q, want %q", test.input, result, test.expected)
		}
	}
}
