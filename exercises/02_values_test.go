package exercises

import (
	"testing"
)

// Exercise: Values and Variables
// Task: Declare and initialize variables of different types
func GetBasicValues() (string, int, bool, float64) {
	// TODO: Declare a string variable with value "go"
	var str string
	
	// TODO: Declare an int variable with value 42
	var num int
	
	// TODO: Declare a bool variable with value true
	var flag bool
	
	// TODO: Declare a float64 variable with value 3.14
	var pi float64
	
	return str, num, flag, pi
}

func TestBasicValues(t *testing.T) {
	str, num, flag, pi := GetBasicValues()
	
	if str != "go" {
		t.Errorf("String value = %q, want %q", str, "go")
	}
	if num != 42 {
		t.Errorf("Int value = %d, want %d", num, 42)
	}
	if flag != true {
		t.Errorf("Bool value = %t, want %t", flag, true)
	}
	if pi != 3.14 {
		t.Errorf("Float value = %f, want %f", pi, 3.14)
	}
}

// Exercise: Constants
// Task: Define constants and use them
func GetConstants() (string, int) {
	// TODO: Define a constant string "constant"
	// TODO: Define a constant int 500000000
	// TODO: Return both constants
	return "", 0
}

func TestConstants(t *testing.T) {
	str, num := GetConstants()
	
	if str != "constant" {
		t.Errorf("String constant = %q, want %q", str, "constant")
	}
	if num != 500000000 {
		t.Errorf("Int constant = %d, want %d", num, 500000000)
	}
}
