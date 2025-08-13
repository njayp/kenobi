package exercises

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Exercise: JSON
// Task: Work with JSON encoding and decoding

// PersonJSON struct for JSON operations
type PersonJSON struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Address *Address `json:"address,omitempty"`
}

// Address struct
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
}

// EncodeJSON: Convert struct to JSON string
func EncodeJSON(person PersonJSON) (string, error) {
	// TODO: Marshal person to JSON and return as string
	return "", nil
}

// DecodeJSON: Convert JSON string to struct
func DecodeJSON(jsonStr string) (PersonJSON, error) {
	// TODO: Unmarshal JSON string to PersonJSON struct
	return PersonJSON{}, nil
}

// EncodeSlice: Convert slice of structs to JSON
func EncodeSlice(people []PersonJSON) (string, error) {
	// TODO: Marshal slice to JSON
	return "", nil
}

// DecodeSlice: Convert JSON array to slice of structs
func DecodeSlice(jsonStr string) ([]PersonJSON, error) {
	// TODO: Unmarshal JSON array to slice
	return nil, nil
}

// EncodeMap: Convert map to JSON
func EncodeMap(data map[string]interface{}) (string, error) {
	// TODO: Marshal map to JSON
	return "", nil
}

// DecodeMap: Convert JSON to map
func DecodeMap(jsonStr string) (map[string]interface{}, error) {
	// TODO: Unmarshal JSON to map
	return nil, nil
}

// Custom JSON marshaling
type Temperature struct {
	Celsius float64
}

// MarshalJSON: Custom JSON marshaling for Temperature
func (t Temperature) MarshalJSON() ([]byte, error) {
	// TODO: Return JSON with both celsius and fahrenheit
	// Format: {"celsius": 25.0, "fahrenheit": 77.0}
	return nil, nil
}

// UnmarshalJSON: Custom JSON unmarshaling for Temperature
func (t *Temperature) UnmarshalJSON(data []byte) error {
	// TODO: Parse JSON and set Celsius field
	// Accept either {"celsius": 25.0} or {"fahrenheit": 77.0}
	return nil
}

func TestEncodeJSON(t *testing.T) {
	person := PersonJSON{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	
	result, err := EncodeJSON(person)
	if err != nil {
		t.Errorf("EncodeJSON returned error: %v", err)
		return
	}
	
	// Parse back to verify structure
	var parsed PersonJSON
	err = json.Unmarshal([]byte(result), &parsed)
	if err != nil {
		t.Errorf("Result is not valid JSON: %v", err)
		return
	}
	
	if parsed.Name != person.Name {
		t.Errorf("Name = %q, want %q", parsed.Name, person.Name)
	}
	if parsed.Age != person.Age {
		t.Errorf("Age = %d, want %d", parsed.Age, person.Age)
	}
}

func TestDecodeJSON(t *testing.T) {
	jsonStr := `{"name":"Bob","age":25,"email":"bob@example.com"}`
	
	result, err := DecodeJSON(jsonStr)
	if err != nil {
		t.Errorf("DecodeJSON returned error: %v", err)
		return
	}
	
	expected := PersonJSON{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
	}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DecodeJSON = %+v, want %+v", result, expected)
	}
}

func TestEncodeSlice(t *testing.T) {
	people := []PersonJSON{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
	}
	
	result, err := EncodeSlice(people)
	if err != nil {
		t.Errorf("EncodeSlice returned error: %v", err)
		return
	}
	
	// Parse back to verify
	var parsed []PersonJSON
	err = json.Unmarshal([]byte(result), &parsed)
	if err != nil {
		t.Errorf("Result is not valid JSON: %v", err)
		return
	}
	
	if len(parsed) != len(people) {
		t.Errorf("Length = %d, want %d", len(parsed), len(people))
	}
}

func TestDecodeSlice(t *testing.T) {
	jsonStr := `[{"name":"Alice","age":30,"email":"alice@example.com"},{"name":"Bob","age":25,"email":"bob@example.com"}]`
	
	result, err := DecodeSlice(jsonStr)
	if err != nil {
		t.Errorf("DecodeSlice returned error: %v", err)
		return
	}
	
	if len(result) != 2 {
		t.Errorf("Length = %d, want 2", len(result))
		return
	}
	
	if result[0].Name != "Alice" {
		t.Errorf("First person name = %q, want %q", result[0].Name, "Alice")
	}
}

func TestEncodeMap(t *testing.T) {
	data := map[string]interface{}{
		"name":   "Charlie",
		"age":    35,
		"active": true,
	}
	
	result, err := EncodeMap(data)
	if err != nil {
		t.Errorf("EncodeMap returned error: %v", err)
		return
	}
	
	// Parse back to verify
	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(result), &parsed)
	if err != nil {
		t.Errorf("Result is not valid JSON: %v", err)
	}
}

func TestDecodeMap(t *testing.T) {
	jsonStr := `{"name":"David","age":40,"active":false}`
	
	result, err := DecodeMap(jsonStr)
	if err != nil {
		t.Errorf("DecodeMap returned error: %v", err)
		return
	}
	
	if result["name"] != "David" {
		t.Errorf("Name = %v, want %q", result["name"], "David")
	}
	// Note: JSON numbers are float64 by default
	if result["age"] != float64(40) {
		t.Errorf("Age = %v, want %v", result["age"], float64(40))
	}
}

func TestCustomMarshalJSON(t *testing.T) {
	temp := Temperature{Celsius: 25.0}
	
	data, err := json.Marshal(temp)
	if err != nil {
		t.Errorf("Marshal returned error: %v", err)
		return
	}
	
	// Should contain both celsius and fahrenheit
	var result map[string]float64
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Errorf("Result is not valid JSON: %v", err)
		return
	}
	
	if result["celsius"] != 25.0 {
		t.Errorf("Celsius = %f, want 25.0", result["celsius"])
	}
	if result["fahrenheit"] != 77.0 {
		t.Errorf("Fahrenheit = %f, want 77.0", result["fahrenheit"])
	}
}

func TestCustomUnmarshalJSON(t *testing.T) {
	// Test unmarshaling from celsius
	celsiusJSON := `{"celsius": 20.0}`
	var temp1 Temperature
	err := json.Unmarshal([]byte(celsiusJSON), &temp1)
	if err != nil {
		t.Errorf("Unmarshal celsius returned error: %v", err)
		return
	}
	if temp1.Celsius != 20.0 {
		t.Errorf("Celsius = %f, want 20.0", temp1.Celsius)
	}
	
	// Test unmarshaling from fahrenheit
	fahrenheitJSON := `{"fahrenheit": 86.0}`
	var temp2 Temperature
	err = json.Unmarshal([]byte(fahrenheitJSON), &temp2)
	if err != nil {
		t.Errorf("Unmarshal fahrenheit returned error: %v", err)
		return
	}
	// 86°F = 30°C
	if temp2.Celsius != 30.0 {
		t.Errorf("Celsius = %f, want 30.0", temp2.Celsius)
	}
}
