package exercises

import (
	"testing"
)

// Exercise: Switch Statements
// Task: Implement switch statements for different scenarios

// BasicSwitch: Return day name based on day number (1=Monday, 7=Sunday)
func BasicSwitch(dayNum int) string {
	// TODO: Use switch to return day name
	// 1="Monday", 2="Tuesday", ..., 7="Sunday", default="Invalid"
	return ""
}

// SwitchMultipleExpressions: Categorize time of day
func SwitchMultipleExpressions(hour int) string {
	// TODO: Use switch with multiple expressions
	// 6,7,8="morning", 12,13="lunch", 18,19,20="evening", default="other"
	return ""
}

// SwitchWithoutExpression: Use switch without expression (like if-else chain)
func SwitchWithoutExpression(score int) string {
	// TODO: Use switch without expression
	// score >= 90: "A", score >= 80: "B", score >= 70: "C", default: "F"
	return ""
}

func TestBasicSwitch(t *testing.T) {
	tests := []struct {
		day      int
		expected string
	}{
		{1, "Monday"},
		{2, "Tuesday"},
		{3, "Wednesday"},
		{4, "Thursday"},
		{5, "Friday"},
		{6, "Saturday"},
		{7, "Sunday"},
		{8, "Invalid"},
		{0, "Invalid"},
	}
	
	for _, test := range tests {
		result := BasicSwitch(test.day)
		if result != test.expected {
			t.Errorf("BasicSwitch(%d) = %q, want %q", test.day, result, test.expected)
		}
	}
}

func TestSwitchMultipleExpressions(t *testing.T) {
	tests := []struct {
		hour     int
		expected string
	}{
		{6, "morning"},
		{7, "morning"},
		{8, "morning"},
		{12, "lunch"},
		{13, "lunch"},
		{18, "evening"},
		{19, "evening"},
		{20, "evening"},
		{15, "other"},
		{22, "other"},
	}
	
	for _, test := range tests {
		result := SwitchMultipleExpressions(test.hour)
		if result != test.expected {
			t.Errorf("SwitchMultipleExpressions(%d) = %q, want %q", test.hour, result, test.expected)
		}
	}
}

func TestSwitchWithoutExpression(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{95, "A"},
		{90, "A"},
		{85, "B"},
		{80, "B"},
		{75, "C"},
		{70, "C"},
		{65, "F"},
		{50, "F"},
	}
	
	for _, test := range tests {
		result := SwitchWithoutExpression(test.score)
		if result != test.expected {
			t.Errorf("SwitchWithoutExpression(%d) = %q, want %q", test.score, result, test.expected)
		}
	}
}
