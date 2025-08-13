package exercises

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Exercise: String Functions
// Task: Work with string manipulation and formatting

// StringBasics: Basic string operations
func StringBasics(str string) (int, string, string) {
	// TODO: Return length, uppercase version, and lowercase version
	return 0, "", ""
}

// StringContains: Check if string contains substring
func StringContains(str, substr string) bool {
	// TODO: Return true if str contains substr
	return false
}

// StringSplit: Split string by delimiter
func StringSplit(str, delimiter string) []string {
	// TODO: Split str by delimiter and return slice of strings
	return nil
}

// StringJoin: Join slice of strings with separator
func StringJoin(strs []string, separator string) string {
	// TODO: Join slice with separator
	return ""
}

// StringReplace: Replace occurrences in string
func StringReplace(str, old, new string) string {
	// TODO: Replace all occurrences of 'old' with 'new' in str
	return ""
}

// StringTrim: Trim whitespace from string
func StringTrim(str string) string {
	// TODO: Remove leading and trailing whitespace
	return ""
}

// StringFormatting: Use fmt.Sprintf for string formatting
func StringFormatting(name string, age int, balance float64) string {
	// TODO: Return formatted string: "Name: <name>, Age: <age>, Balance: $<balance:.2f>"
	return ""
}

// StringParsing: Parse strings to different types
func StringParsing(intStr, floatStr, boolStr string) (int, float64, bool, error) {
	// TODO: Parse each string to appropriate type
	// TODO: Return parsed values and error if any parsing fails
	return 0, 0.0, false, nil
}

// StringValidation: Validate string format
func StringValidation(email string) bool {
	// TODO: Basic email validation (must contain @ and . after @)
	return false
}

// StringIndexing: Find index of substring
func StringIndexing(str, substr string) int {
	// TODO: Return index of first occurrence of substr in str, -1 if not found
	return -1
}

func TestStringBasics(t *testing.T) {
	length, upper, lower := StringBasics("Hello World")
	
	if length != 11 {
		t.Errorf("Length = %d, want 11", length)
	}
	if upper != "HELLO WORLD" {
		t.Errorf("Upper = %q, want %q", upper, "HELLO WORLD")
	}
	if lower != "hello world" {
		t.Errorf("Lower = %q, want %q", lower, "hello world")
	}
}

func TestStringContains(t *testing.T) {
	tests := []struct {
		str, substr string
		expected    bool
	}{
		{"hello world", "world", true},
		{"hello world", "foo", false},
		{"Go programming", "Go", true},
		{"case sensitive", "CASE", false},
	}
	
	for _, test := range tests {
		result := StringContains(test.str, test.substr)
		if result != test.expected {
			t.Errorf("StringContains(%q, %q) = %t, want %t", 
				test.str, test.substr, result, test.expected)
		}
	}
}

func TestStringSplit(t *testing.T) {
	result := StringSplit("apple,banana,cherry", ",")
	expected := []string{"apple", "banana", "cherry"}
	
	if len(result) != len(expected) {
		t.Errorf("Length = %d, want %d", len(result), len(expected))
		return
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("Index %d: got %q, want %q", i, result[i], exp)
		}
	}
}

func TestStringJoin(t *testing.T) {
	strs := []string{"apple", "banana", "cherry"}
	result := StringJoin(strs, ", ")
	expected := "apple, banana, cherry"
	
	if result != expected {
		t.Errorf("StringJoin = %q, want %q", result, expected)
	}
}

func TestStringReplace(t *testing.T) {
	result := StringReplace("hello world hello", "hello", "hi")
	expected := "hi world hi"
	
	if result != expected {
		t.Errorf("StringReplace = %q, want %q", result, expected)
	}
}

func TestStringTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello  ", "hello"},
		{"\t\nworld\t\n", "world"},
		{"no spaces", "no spaces"},
		{"   ", ""},
	}
	
	for _, test := range tests {
		result := StringTrim(test.input)
		if result != test.expected {
			t.Errorf("StringTrim(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestStringFormatting(t *testing.T) {
	result := StringFormatting("Alice", 30, 1234.567)
	expected := "Name: Alice, Age: 30, Balance: $1234.57"
	
	if result != expected {
		t.Errorf("StringFormatting = %q, want %q", result, expected)
	}
}

func TestStringParsing(t *testing.T) {
	intVal, floatVal, boolVal, err := StringParsing("42", "3.14", "true")
	
	if err != nil {
		t.Errorf("StringParsing returned error: %v", err)
		return
	}
	
	if intVal != 42 {
		t.Errorf("Int = %d, want 42", intVal)
	}
	if floatVal != 3.14 {
		t.Errorf("Float = %f, want 3.14", floatVal)
	}
	if boolVal != true {
		t.Errorf("Bool = %t, want true", boolVal)
	}
	
	// Test error case
	_, _, _, err = StringParsing("not_a_number", "3.14", "true")
	if err == nil {
		t.Error("StringParsing should return error for invalid int")
	}
}

func TestStringValidation(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"user@example.com", true},
		{"test.email@domain.org", true},
		{"invalid_email", false},
		{"no_at_symbol.com", false},
		{"user@no_dot", false},
		{"@example.com", false},
	}
	
	for _, test := range tests {
		result := StringValidation(test.email)
		if result != test.expected {
			t.Errorf("StringValidation(%q) = %t, want %t", 
				test.email, result, test.expected)
		}
	}
}

func TestStringIndexing(t *testing.T) {
	tests := []struct {
		str, substr string
		expected    int
	}{
		{"hello world", "world", 6},
		{"hello world", "hello", 0},
		{"hello world", "foo", -1},
		{"programming", "gram", 3},
	}
	
	for _, test := range tests {
		result := StringIndexing(test.str, test.substr)
		if result != test.expected {
			t.Errorf("StringIndexing(%q, %q) = %d, want %d", 
				test.str, test.substr, result, test.expected)
		}
	}
}

// Advanced string exercises
func TestStringBuilder(t *testing.T) {
	// TODO: Students should implement using strings.Builder for efficient string building
}

func BuildLargeString(words []string) string {
	// TODO: Use strings.Builder to efficiently concatenate many strings
	return ""
}

func TestStringBuilder(t *testing.T) {
	words := []string{"Go", "is", "an", "awesome", "language"}
	result := BuildLargeString(words)
	expected := "Go is an awesome language"
	
	if result != expected {
		t.Errorf("BuildLargeString = %q, want %q", result, expected)
	}
}
