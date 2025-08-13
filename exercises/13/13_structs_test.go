package exercises

import (
	"testing"
)

// Exercise: Structs
// Task: Define and work with structs

// Person represents a person with name and age
type Person struct {
	// TODO: Add Name field of type string
	// TODO: Add Age field of type int
}

// NewPerson: Constructor function for Person
func NewPerson(name string, age int) Person {
	// TODO: Return a Person with given name and age
	return Person{}
}

// GetInfo: Method to get person info as string
func (p Person) GetInfo() string {
	// TODO: Return string in format "Name: <name>, Age: <age>"
	return ""
}

// HaveBirthday: Method to increment age (pointer receiver needed)
func (p *Person) HaveBirthday() {
	// TODO: Increment the person's age by 1
}

// Rectangle represents a rectangle with width and height
type Rectangle struct {
	// TODO: Add Width and Height fields of type float64
}

// Area: Calculate area of rectangle
func (r Rectangle) Area() float64 {
	// TODO: Return width * height
	return 0
}

// Scale: Scale the rectangle by a factor (pointer receiver)
func (r *Rectangle) Scale(factor float64) {
	// TODO: Multiply both width and height by factor
}

func TestPersonStruct(t *testing.T) {
	// Test struct creation
	person := NewPerson("Alice", 25)

	if person.Name != "Alice" {
		t.Errorf("Name = %q, want %q", person.Name, "Alice")
	}
	if person.Age != 25 {
		t.Errorf("Age = %d, want %d", person.Age, 25)
	}
}

func TestGetInfo(t *testing.T) {
	person := NewPerson("Bob", 30)
	info := person.GetInfo()
	expected := "Name: Bob, Age: 30"

	if info != expected {
		t.Errorf("GetInfo() = %q, want %q", info, expected)
	}
}

func TestHaveBirthday(t *testing.T) {
	person := NewPerson("Charlie", 20)
	person.HaveBirthday()

	if person.Age != 21 {
		t.Errorf("After birthday, Age = %d, want 21", person.Age)
	}
}

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	area := rect.Area()
	expected := 15.0

	if area != expected {
		t.Errorf("Area() = %f, want %f", area, expected)
	}
}

func TestRectangleScale(t *testing.T) {
	rect := Rectangle{Width: 4.0, Height: 2.0}
	rect.Scale(2.0)

	if rect.Width != 8.0 {
		t.Errorf("After scaling, Width = %f, want 8.0", rect.Width)
	}
	if rect.Height != 4.0 {
		t.Errorf("After scaling, Height = %f, want 4.0", rect.Height)
	}
}

// Test struct embedding
type Employee struct {
	Person // TODO: Embed Person struct
	// TODO: Add Salary field of type int
}

func TestStructEmbedding(t *testing.T) {
	emp := Employee{
		Person: Person{Name: "Dave", Age: 35},
		Salary: 50000,
	}

	// Should be able to access embedded fields directly
	if emp.Name != "Dave" {
		t.Errorf("Embedded Name = %q, want %q", emp.Name, "Dave")
	}
	if emp.Age != 35 {
		t.Errorf("Embedded Age = %d, want %d", emp.Age, 35)
	}
	if emp.Salary != 50000 {
		t.Errorf("Salary = %d, want %d", emp.Salary, 50000)
	}
}
