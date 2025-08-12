package exercises

import (
	"errors"
	"strconv"
	"testing"
)

// Exercise: Error Handling
// Task: Work with errors in Go

// DivideNumbers: Divide two numbers, return error if dividing by zero
func DivideNumbers(a, b float64) (float64, error) {
	// TODO: Return error if b is 0, otherwise return a/b
	return 0, nil
}

// ParseAndAdd: Parse two strings as integers and add them
func ParseAndAdd(s1, s2 string) (int, error) {
	// TODO: Parse both strings as integers
	// TODO: Return their sum, or return error if parsing fails
	return 0, nil
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

// Error method to implement error interface
func (e *ValidationError) Error() string {
	// TODO: Return formatted error message: "Field: Message"
	return ""
}

// ValidateAge: Validate age and return custom error if invalid
func ValidateAge(age int) error {
	// TODO: Return ValidationError if age < 0 or age > 150
	// TODO: Use "age" as field and appropriate message
	return nil
}

// ChainedErrors: Demonstrate error handling chain
func ChainedErrors(s string) (int, error) {
	// TODO: Parse string as int
	// TODO: If parsing fails, return 0 and the parse error
	// TODO: If result is negative, return ValidationError
	// TODO: Otherwise return the parsed value
	return 0, nil
}

func TestDivideNumbers(t *testing.T) {
	// Test successful division
	result, err := DivideNumbers(10, 2)
	if err != nil {
		t.Errorf("DivideNumbers(10, 2) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("DivideNumbers(10, 2) = %f, want 5", result)
	}
	
	// Test division by zero
	_, err = DivideNumbers(10, 0)
	if err == nil {
		t.Error("DivideNumbers(10, 0) should return error")
	}
}

func TestParseAndAdd(t *testing.T) {
	// Test successful parsing and addition
	result, err := ParseAndAdd("5", "3")
	if err != nil {
		t.Errorf("ParseAndAdd(\"5\", \"3\") returned error: %v", err)
	}
	if result != 8 {
		t.Errorf("ParseAndAdd(\"5\", \"3\") = %d, want 8", result)
	}
	
	// Test parsing error
	_, err = ParseAndAdd("abc", "3")
	if err == nil {
		t.Error("ParseAndAdd(\"abc\", \"3\") should return error")
	}
	
	_, err = ParseAndAdd("5", "xyz")
	if err == nil {
		t.Error("ParseAndAdd(\"5\", \"xyz\") should return error")
	}
}

func TestValidationError(t *testing.T) {
	err := &ValidationError{
		Field:   "username",
		Message: "cannot be empty",
	}
	
	expected := "username: cannot be empty"
	if err.Error() != expected {
		t.Errorf("ValidationError.Error() = %q, want %q", err.Error(), expected)
	}
}

func TestValidateAge(t *testing.T) {
	// Test valid age
	err := ValidateAge(25)
	if err != nil {
		t.Errorf("ValidateAge(25) returned error: %v", err)
	}
	
	// Test negative age
	err = ValidateAge(-5)
	if err == nil {
		t.Error("ValidateAge(-5) should return error")
	}
	
	// Check if it's a ValidationError
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Error("ValidateAge(-5) should return ValidationError")
	}
	
	// Test age too high
	err = ValidateAge(200)
	if err == nil {
		t.Error("ValidateAge(200) should return error")
	}
}

func TestChainedErrors(t *testing.T) {
	// Test successful parsing of positive number
	result, err := ChainedErrors("42")
	if err != nil {
		t.Errorf("ChainedErrors(\"42\") returned error: %v", err)
	}
	if result != 42 {
		t.Errorf("ChainedErrors(\"42\") = %d, want 42", result)
	}
	
	// Test parsing error
	_, err = ChainedErrors("abc")
	if err == nil {
		t.Error("ChainedErrors(\"abc\") should return error")
	}
	
	// Should be a strconv.NumError (from parsing)
	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Error("ChainedErrors(\"abc\") should return strconv.NumError")
	}
	
	// Test negative number (should return ValidationError)
	_, err = ChainedErrors("-5")
	if err == nil {
		t.Error("ChainedErrors(\"-5\") should return error")
	}
	
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Error("ChainedErrors(\"-5\") should return ValidationError")
	}
}

// Test error wrapping (Go 1.13+)
func TestErrorWrapping(t *testing.T) {
	// This is for students to learn about errors.Is and errors.As
	originalErr := errors.New("original error")
	wrappedErr := errors.New("wrapped: " + originalErr.Error())
	
	// TODO: Students should implement error wrapping properly
	// This test should be updated to use fmt.Errorf with %w verb
	_ = wrappedErr
}
