package exercises

import (
	"regexp"
	"testing"
)

// Exercise: Regular Expressions
// Task: Work with regular expressions for pattern matching

// BasicMatch: Check if string matches a pattern
func BasicMatch(pattern, text string) (bool, error) {
	// TODO: Use regexp.MatchString to check if text matches pattern
	return false, nil
}

// CompileAndMatch: Compile regex and use it for matching
func CompileAndMatch(pattern, text string) (bool, error) {
	// TODO: Compile regex pattern and use it to match text
	return false, nil
}

// FindMatches: Find all matches of pattern in text
func FindMatches(pattern, text string) ([]string, error) {
	// TODO: Find all matches of pattern in text
	return nil, nil
}

// FindSubmatch: Find first match and capture groups
func FindSubmatch(pattern, text string) ([]string, error) {
	// TODO: Find first match and return all capture groups
	return nil, nil
}

// ReplacePattern: Replace all matches with replacement
func ReplacePattern(pattern, text, replacement string) (string, error) {
	// TODO: Replace all matches of pattern with replacement
	return "", nil
}

// ValidateEmail: Use regex to validate email format
func ValidateEmail(email string) bool {
	// TODO: Use regex to validate email format
	// Pattern should match: word characters, @, word characters, ., word characters
	return false
}

// ExtractNumbers: Extract all numbers from text
func ExtractNumbers(text string) []string {
	// TODO: Use regex to find all numbers (sequences of digits) in text
	return nil
}

// SplitByPattern: Split text by regex pattern
func SplitByPattern(pattern, text string) ([]string, error) {
	// TODO: Split text using regex pattern as delimiter
	return nil, nil
}

func TestBasicMatch(t *testing.T) {
	tests := []struct {
		pattern, text string
		expected      bool
	}{
		{"hello", "hello world", true},
		{"world", "hello world", true},
		{"foo", "hello world", false},
		{"[0-9]+", "abc123def", true},
		{"[0-9]+", "abcdef", false},
	}
	
	for _, test := range tests {
		result, err := BasicMatch(test.pattern, test.text)
		if err != nil {
			t.Errorf("BasicMatch(%q, %q) returned error: %v", test.pattern, test.text, err)
			continue
		}
		if result != test.expected {
			t.Errorf("BasicMatch(%q, %q) = %t, want %t", 
				test.pattern, test.text, result, test.expected)
		}
	}
}

func TestCompileAndMatch(t *testing.T) {
	result, err := CompileAndMatch("^[a-z]+$", "hello")
	if err != nil {
		t.Errorf("CompileAndMatch returned error: %v", err)
		return
	}
	if !result {
		t.Error("CompileAndMatch should return true for lowercase letters")
	}
	
	result, err = CompileAndMatch("^[a-z]+$", "Hello")
	if err != nil {
		t.Errorf("CompileAndMatch returned error: %v", err)
		return
	}
	if result {
		t.Error("CompileAndMatch should return false for mixed case")
	}
}

func TestFindMatches(t *testing.T) {
	matches, err := FindMatches("[0-9]+", "abc123def456ghi")
	if err != nil {
		t.Errorf("FindMatches returned error: %v", err)
		return
	}
	
	expected := []string{"123", "456"}
	if len(matches) != len(expected) {
		t.Errorf("Found %d matches, want %d", len(matches), len(expected))
		return
	}
	
	for i, exp := range expected {
		if matches[i] != exp {
			t.Errorf("Match %d: got %q, want %q", i, matches[i], exp)
		}
	}
}

func TestFindSubmatch(t *testing.T) {
	// Test email pattern with capture groups
	groups, err := FindSubmatch("([a-z]+)@([a-z]+)\\.([a-z]+)", "test@example.com")
	if err != nil {
		t.Errorf("FindSubmatch returned error: %v", err)
		return
	}
	
	// Should return: full match, username, domain, tld
	expected := []string{"test@example.com", "test", "example", "com"}
	if len(groups) != len(expected) {
		t.Errorf("Found %d groups, want %d", len(groups), len(expected))
		return
	}
	
	for i, exp := range expected {
		if groups[i] != exp {
			t.Errorf("Group %d: got %q, want %q", i, groups[i], exp)
		}
	}
}

func TestReplacePattern(t *testing.T) {
	result, err := ReplacePattern("[0-9]+", "abc123def456", "X")
	if err != nil {
		t.Errorf("ReplacePattern returned error: %v", err)
		return
	}
	
	expected := "abcXdefX"
	if result != expected {
		t.Errorf("ReplacePattern = %q, want %q", result, expected)
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"user@example.com", true},
		{"test@domain.org", true},
		{"invalid.email", false},
		{"user@", false},
		{"@domain.com", false},
		{"user@domain", false},
		{"user@domain.", false},
	}
	
	for _, test := range tests {
		result := ValidateEmail(test.email)
		if result != test.expected {
			t.Errorf("ValidateEmail(%q) = %t, want %t", 
				test.email, result, test.expected)
		}
	}
}

func TestExtractNumbers(t *testing.T) {
	numbers := ExtractNumbers("I have 5 apples and 10 oranges, total 15 fruits")
	expected := []string{"5", "10", "15"}
	
	if len(numbers) != len(expected) {
		t.Errorf("Found %d numbers, want %d", len(numbers), len(expected))
		return
	}
	
	for i, exp := range expected {
		if numbers[i] != exp {
			t.Errorf("Number %d: got %q, want %q", i, numbers[i], exp)
		}
	}
}

func TestSplitByPattern(t *testing.T) {
	parts, err := SplitByPattern("[,;]", "apple,banana;cherry,date")
	if err != nil {
		t.Errorf("SplitByPattern returned error: %v", err)
		return
	}
	
	expected := []string{"apple", "banana", "cherry", "date"}
	if len(parts) != len(expected) {
		t.Errorf("Split into %d parts, want %d", len(parts), len(expected))
		return
	}
	
	for i, exp := range expected {
		if parts[i] != exp {
			t.Errorf("Part %d: got %q, want %q", i, parts[i], exp)
		}
	}
}

// Advanced regex exercises
func TestNamedGroups(t *testing.T) {
	// TODO: Students should implement named capture groups
}

func ParseLogEntry(logLine string) (map[string]string, error) {
	// TODO: Parse log entry with named capture groups
	// Example: "2024-01-15 10:30:45 INFO Application started"
	// Should extract: date, time, level, message
	return nil, nil
}

func TestParseLogEntry(t *testing.T) {
	result, err := ParseLogEntry("2024-01-15 10:30:45 INFO Application started")
	if err != nil {
		t.Errorf("ParseLogEntry returned error: %v", err)
		return
	}
	
	expected := map[string]string{
		"date":    "2024-01-15",
		"time":    "10:30:45",
		"level":   "INFO",
		"message": "Application started",
	}
	
	for key, expectedValue := range expected {
		if result[key] != expectedValue {
			t.Errorf("Key %q: got %q, want %q", key, result[key], expectedValue)
		}
	}
}

func TestComplexValidation(t *testing.T) {
	// TODO: Students should implement complex validation patterns
}

func ValidatePassword(password string) bool {
	// TODO: Validate password requirements:
	// - At least 8 characters
	// - Contains at least one uppercase letter
	// - Contains at least one lowercase letter  
	// - Contains at least one digit
	// - Contains at least one special character (!@#$%^&*)
	return false
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"Password123!", true},
		{"weakpass", false},        // no uppercase, digit, special
		{"PASSWORD123!", false},    // no lowercase
		{"Password!", false},       // no digit
		{"Password123", false},     // no special character
		{"Pass1!", false},          // too short
	}
	
	for _, test := range tests {
		result := ValidatePassword(test.password)
		if result != test.expected {
			t.Errorf("ValidatePassword(%q) = %t, want %t", 
				test.password, result, test.expected)
		}
	}
}
