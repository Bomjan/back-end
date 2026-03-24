package main

import (
	"encoding/json"
	"fmt"
)

// Spouse represents a person's spouse information.
// Used as a nested type within the Person struct.
type Spouse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Person represents a person with their personal information,
// including marital status, spouse details, and children.
type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Married  bool     `json:"married"`
	Spouse   Spouse   `json:"spouse"`
	Children []string `json:"children"`
}

// Employee represents an employee record with personal and work information.
// The JSON struct tags show: custom naming (full_name), omitempty behavior,
// field exclusion (-), and optional fields.
type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"full_name"`        // Rename field in JSON
	Email    string `json:"email,omitempty"`  // Omit if empty
	Password string `json:"-"`                // Always exclude from JSON
	Salary   int    `json:"salary,omitempty"` // Omit if zero value
	Active   bool   `json:"is_active"`        // Rename field
}

// APIResponse is a standard HTTP API response wrapper.
// Status indicates success or failure ("success" or "error").
// Data contains the response payload and is omitted if empty.
// Error contains error message details and is omitted on success.
type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func main() {
	fmt.Println("========== JSON DECODING (Unmarshal) ==========")
	example1Unmarshal()

	fmt.Println("\n========== JSON ENCODING (Marshal) ==========")
	example2Marshal()

	fmt.Println("\n========== STRUCT TAGS & OMITEMPTY ==========")
	example3Tags()

	fmt.Println("\n========== PRETTY PRINTING JSON ==========")
	example4PrettyPrint()

	fmt.Println("\n========== DYNAMIC DATA WITH MAPS ==========")
	example5Maps()

	fmt.Println("\n========== ERROR HANDLING ==========")
	example6Errors()

	fmt.Println("\n========== API RESPONSE PATTERN ==========")
	example7APIResponse()
}

// Example 1: Decoding JSON bytes to Go struct
func example1Unmarshal() {
	fmt.Println("Example 1: Decoding JSON → Go Struct (Unmarshal)")
	fmt.Println("---")

	jsonData := []byte(`
		{
			"name":"Sundra Bomjan",
			"age":90,
			"married":false,
			"spouse":{
				"name":"IDK",
				"age":10000
			},
			"children":["Su", "Bo"]
		}
	`)

	var p Person

	err := json.Unmarshal(jsonData, &p)

	if err != nil {
		fmt.Println("Error Found:", err)
		return
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Married: %v\n", p.Married)
	fmt.Printf("Spouse: %s (Age: %d)\n", p.Spouse.Name, p.Spouse.Age)
	fmt.Printf("Children: %v\n", p.Children)
}

// Example 2: Encoding Go struct to JSON bytes
func example2Marshal() {
	fmt.Println("Example 2: Encoding Go Struct → JSON (Marshal)")
	fmt.Println("---")

	person := Person{
		Name:    "Alice Johnson",
		Age:     28,
		Married: true,
		Spouse: Spouse{
			Name: "Bob Johnson",
			Age:  30,
		},
		Children: []string{"Charlie", "Diana"},
	}

	// Marshal to compact JSON
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	fmt.Println("Compact JSON:")
	fmt.Println(string(jsonBytes))
	fmt.Println()
}

// Example 3: Struct tags - omitempty, -, custom field names
func example3Tags() {
	fmt.Println("Example 3: Struct Tags (omitempty, -, custom names)")
	fmt.Println("---")

	// Employee with all fields filled
	emp1 := Employee{
		ID:       1,
		Name:     "Eve Smith",
		Email:    "eve@example.com",
		Password: "secret123", // This will be excluded by json:"-"
		Salary:   75000,
		Active:   true,
	}

	// Employee with empty email and zero salary
	emp2 := Employee{
		ID:       2,
		Name:     "Frank Doe",
		Email:    "", // Will be omitted
		Password: "secret456",
		Salary:   0, // Will be omitted (zero value)
		Active:   true,
	}

	fmt.Println("Employee 1 (all fields):")
	json1, err := json.Marshal(emp1)
	if err != nil {
		fmt.Println("Error encoding emp1:", err)
		return
	}
	fmt.Println(string(json1))

	fmt.Println("\nEmployee 2 (empty email and zero salary omitted):")
	json2, err := json.Marshal(emp2)
	if err != nil {
		fmt.Println("Error encoding emp2:", err)
		return
	}
	fmt.Println(string(json2))

	fmt.Println("\nNote: Password field is NOT in JSON (json:\"-\")")
	fmt.Println("Note: Email and Salary are omitted when empty (json:\"...,omitempty\")")
}

// Example 4: Pretty printing JSON with indentation
func example4PrettyPrint() {
	fmt.Println("Example 4: Pretty Printing JSON (MarshalIndent)")
	fmt.Println("---")

	person := Person{
		Name:     "Grace Lee",
		Age:      35,
		Married:  false,
		Spouse:   Spouse{},
		Children: []string{"Henry", "Iris", "Jack"},
	}

	// Compact JSON
	compact, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding compact:", err)
		return
	}
	fmt.Println("Compact JSON:")
	fmt.Println(string(compact))

	// Pretty-printed JSON with 2-space indent
	pretty, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error encoding pretty:", err)
		return
	}
	fmt.Println("\nPretty-Printed JSON:")
	fmt.Println(string(pretty))
}

// Example 5: Dynamic data with maps
func example5Maps() {
	fmt.Println("Example 5: Dynamic Data with Maps")
	fmt.Println("---")

	// Unmarshal into a map (unknown structure)
	jsonString := `{"name":"Kevin","role":"Engineer","experience":5,"languages":["Go","Python","JavaScript"]}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonString), &data)

	fmt.Println("Decoded map:")
	for key, value := range data {
		fmt.Printf("  %s: %v (type: %T)\n", key, value, value)
	}

	// Create and marshal a dynamic map
	dynamicData := map[string]interface{}{
		"id":     42,
		"name":   "Luna",
		"active": true,
		"tags":   []string{"admin", "developer"},
	}

	jsonBytes, _ := json.Marshal(dynamicData)
	fmt.Println("\nMarshalled map to JSON:")
	fmt.Println(string(jsonBytes))
}

// Example 6: Error handling
func example6Errors() {
	fmt.Println("Example 6: Error Handling")
	fmt.Println("---")

	// Invalid JSON - missing closing brace
	invalidJSON := []byte(`{"name":"Mike","age":25`)

	var p Person
	err := json.Unmarshal(invalidJSON, &p)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		fmt.Println("Type of error:", fmt.Sprintf("%T", err))
	}

	// Valid JSON but wrong types
	wrongTypes := []byte(`{"name":"Nancy","age":"not a number"}`)
	err2 := json.Unmarshal(wrongTypes, &p)

	if err2 != nil {
		fmt.Printf("\nERROR: %v\n", err2)
	}
}

// Example 7: API response pattern
func example7APIResponse() {
	fmt.Println("Example 7: API Response Pattern")
	fmt.Println("---")

	// Success response
	successResp := APIResponse{
		Status: "success",
		Data: map[string]string{
			"message": "User created successfully",
			"userId":  "12345",
		},
	}

	successJSON, err := json.MarshalIndent(successResp, "", "  ")
	if err != nil {
		fmt.Println("Error encoding success response:", err)
		return
	}
	fmt.Println("Success Response:")
	fmt.Println(string(successJSON))

	// Error response
	errorResp := APIResponse{
		Status: "error",
		Error:  "Invalid email format",
	}

	errorJSON, err := json.MarshalIndent(errorResp, "", "  ")
	if err != nil {
		fmt.Println("Error encoding error response:", err)
		return
	}
	fmt.Println("\nError Response:")
	fmt.Println(string(errorJSON))
}
