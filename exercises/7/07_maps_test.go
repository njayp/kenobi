package exercises

import (
	"reflect"
	"testing"
)

// Exercise: Maps
// Task: Work with maps (hash tables)

// CreateMap: Create and return a map with string keys and int values
func CreateMap() map[string]int {
	// TODO: Create map with entries: "apple"->5, "banana"->3, "orange"->8
	return nil
}

// MapOperations: Perform various map operations
func MapOperations() (int, bool, bool, int) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	
	// TODO: Get value for key "b"
	valueB := 0
	
	// TODO: Check if key "d" exists (should be false)
	_, existsD := false, false
	
	// TODO: Add new entry "d" -> 4
	
	// TODO: Check if key "d" exists now (should be true)
	_, existsD2 := false, false
	
	// TODO: Delete key "a"
	
	// TODO: Return: valueB, existsD, existsD2, length of final map
	return valueB, existsD, existsD2, 0
}

// MakeMap: Use make to create map
func MakeMap() map[string]int {
	// TODO: Use make to create empty map, then add "first"->1, "second"->2
	return nil
}

func TestCreateMap(t *testing.T) {
	result := CreateMap()
	expected := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CreateMap() = %v, want %v", result, expected)
	}
}

func TestMapOperations(t *testing.T) {
	valueB, existsD, existsD2, finalLength := MapOperations()
	
	if valueB != 2 {
		t.Errorf("valueB = %d, want 2", valueB)
	}
	if existsD != false {
		t.Errorf("existsD = %t, want false", existsD)
	}
	if existsD2 != true {
		t.Errorf("existsD2 = %t, want true", existsD2)
	}
	if finalLength != 3 {
		t.Errorf("finalLength = %d, want 3", finalLength)
	}
}

func TestMakeMap(t *testing.T) {
	result := MakeMap()
	expected := map[string]int{
		"first":  1,
		"second": 2,
	}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MakeMap() = %v, want %v", result, expected)
	}
}
