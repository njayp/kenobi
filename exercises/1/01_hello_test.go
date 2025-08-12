package exercises

import (
	"testing"
)

// Exercise: Hello World
// Task: Make the HelloWorld function return "Hello, World!"
func HelloWorld() string {
	// TODO: Return "Hello, World!"
	return ""
}

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := HelloWorld()
	if result != expected {
		t.Errorf("HelloWorld() = %q, want %q", result, expected)
	}
}
