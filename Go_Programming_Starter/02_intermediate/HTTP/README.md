# HTTP, Encoding/Decoding & Data Serialization in Go

This module introduces you to HTTP fundamentals and how to work with data serialization in Go, specifically focusing on JSON encoding and decoding.

## Table of Contents

1. [HTTP Fundamentals](#http-fundamentals)
2. [JSON Encoding & Decoding](#json-encoding--decoding)
3. [Go's `encoding/json` Package](#gos-encodingjson-package)
4. [Data Serialization](#data-serialization)
5. [Common Patterns](#common-patterns)

---

## HTTP Fundamentals

HTTP (HyperText Transfer Protocol) is the foundation of web communication. It's a **request-response protocol** where a client sends a request to a server, and the server responds with data.

### Request Structure

An HTTP request contains:
- **Method**: What action to perform (`GET`, `POST`, `PUT`, `DELETE`, `PATCH`)
- **URL**: Where to send the request
- **Headers**: Metadata about the request (e.g., `Content-Type: application/json`)
- **Body**: Data sent with the request (optional for GET, required for POST)

Example:
```
POST /api/users HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 27

{"name":"John","age":30}
```

### Response Structure

An HTTP response contains:
- **Status Code**: Result of the request (200 = OK, 404 = Not Found, 500 = Server Error)
- **Headers**: Metadata about the response
- **Body**: The actual data returned

Example:
```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 32

{"id":1,"name":"John","age":30}
```

### Common HTTP Methods

| Method | Purpose | Has Body? |
|--------|---------|-----------|
| `GET` | Retrieve data | No |
| `POST` | Create new data | Yes |
| `PUT` | Replace entire resource | Yes |
| `DELETE` | Remove data | No |
| `PATCH` | Partially update data | Yes |

---

## JSON Encoding & Decoding

JSON (JavaScript Object Notation) is a lightweight, human-readable format for exchanging data over HTTP. Go treats JSON as a series of **bytes**, not native objects.

### Understanding the Process

**Encoding (Go → JSON bytes)**
```
Go struct → JSON string → bytes transmitted over HTTP
```

**Decoding (JSON bytes → Go)**
```
HTTP bytes → JSON string → Go struct
```

### Why This Matters

- **Encoding**: When your Go program sends data to a client, you convert your structured data into JSON bytes
- **Decoding**: When your Go program receives data from a client, you convert JSON bytes back into Go structs

---

## Go's `encoding/json` Package

Go provides a built-in `encoding/json` package for JSON serialization.

### Key Functions

#### 1. **`json.Unmarshal()`** - Decode JSON bytes to Go struct

Converts JSON bytes into a Go struct. Think of it as "unpacking" the JSON data.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// JSON data as bytes
	jsonData := []byte(`{"Name":"Alice","Age":30}`)
	
	// Create an empty struct to hold the decoded data
	var p Person
	
	// Unmarshal: convert JSON bytes → Go struct
	err := json.Unmarshal(jsonData, &p)
	
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	
	fmt.Println("Decoded person:", p)
	// Output: Decoded person: {Alice 30}
}
```

**Important Notes:**
- Pass a **pointer** (`&p`) to `Unmarshal()` because it needs to modify the struct
- The struct field names must match the JSON keys exactly (case-sensitive!)
- If a JSON key doesn't match, that field stays at its zero value

#### 2. **`json.Marshal()`** - Encode Go struct to JSON bytes

Converts a Go struct into JSON bytes. This is the opposite of `Unmarshal()`.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Bob", Age: 25}
	
	// Marshal: convert Go struct → JSON bytes
	jsonBytes, err := json.Marshal(p)
	
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}
	
	// Convert bytes to string for printing
	jsonString := string(jsonBytes)
	fmt.Println("Encoded JSON:", jsonString)
	// Output: Encoded JSON: {"Name":"Bob","Age":25}
}
```

---

## Data Serialization

### Structs and Field Tags

When using `json.Marshal()` and `json.Unmarshal()`, Go uses **struct tags** to control JSON behavior.

```go
type Person struct {
	Name    string `json:"name"`           // Use "name" as JSON key
	Age     int    `json:"age"`            // Use "age" as JSON key
	Email   string `json:"email,omitempty"` // Omit if empty
	Secret  string `json:"-"`              // Ignore this field
}
```

**Tag Options:**

| Tag | Meaning |
|-----|---------|
| `json:"fieldName"` | Rename field in JSON |
| `json:"fieldName,omitempty"` | Omit from JSON if zero value |
| `json:"-"` | Always exclude from JSON |
| `json:"fieldName,string"` | Convert to/from string in JSON |

### Example with Tags

```go
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Secret  string `json:"-"`
}

func main() {
	p := Person{Name: "Charlie", Age: 0, Secret: "hidden"}
	
	jsonBytes, _ := json.Marshal(p)
	fmt.Println(string(jsonBytes))
	// Output: {"name":"Charlie"}
	// Note: Age is omitted (zero value), Secret is excluded
}
```

### Nested Structs

JSON can represent nested objects. Go handles this naturally with nested structs:

```go
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func main() {
	jsonData := []byte(`{
		"name":"Diana",
		"address":{"city":"NYC","state":"NY"}
	}`)
	
	var emp Employee
	json.Unmarshal(jsonData, &emp)
	
	fmt.Println(emp.Name, emp.Address.City)
	// Output: Diana NYC
}
```

### Arrays and Slices

JSON arrays map to Go slices:

```go
type Team struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func main() {
	jsonData := []byte(`{
		"name":"Dev Team",
		"members":["Alice","Bob","Charlie"]
	}`)
	
	var team Team
	json.Unmarshal(jsonData, &team)
	
	fmt.Println(team.Members)
	// Output: [Alice Bob Charlie]
}
```

---

## Common Patterns

### Pattern 1: API Response Wrapper

```go
type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success response
response := APIResponse{Status: "success", Data: myData}
jsonBytes, _ := json.Marshal(response)

// Error response
errorResponse := APIResponse{Status: "error", Error: "Not found"}
jsonBytes, _ := json.Marshal(errorResponse)
```

### Pattern 2: Flexible Data with `map[string]interface{}`

When you don't know the JSON structure in advance:

```go
var data map[string]interface{}

jsonString := `{"name":"Eve","age":28,"active":true}`
json.Unmarshal([]byte(jsonString), &data)

// Access dynamically
fmt.Println(data["name"])   // Eve
fmt.Println(data["age"])    // 28
fmt.Println(data["active"]) // true
```

### Pattern 3: Pretty Printing JSON

```go
data := map[string]string{"name": "Frank", "role": "Engineer"}

// Compact JSON (default)
compact, _ := json.Marshal(data)
fmt.Println(string(compact))
// Output: {"name":"Frank","role":"Engineer"}

// Pretty-printed JSON
pretty, _ := json.MarshalIndent(data, "", "  ")
fmt.Println(string(pretty))
// Output:
// {
//   "name": "Frank",
//   "role": "Engineer"
// }
```

---

## JavaScript Developer's Perspective

If you're coming from JavaScript, here's the mapping:

| Concept | JavaScript | Go |
|---------|------------|-----|
| Object | `{name: "John"}` | `struct` with tags |
| Array | `["a", "b"]` | `[]string` slice |
| Parse JSON | `JSON.parse()` | `json.Unmarshal()` |
| Stringify | `JSON.stringify()` | `json.Marshal()` |
| Dynamic object | `{}` or `map` | `map[string]interface{}` |

### Key Differences

1. **Go is Typed**: A JavaScript object can have any fields. A Go struct has predefined fields.
2. **Strict Matching**: JSON keys must exactly match struct field names (case matters!)
3. **Error Handling**: JavaScript silently ignores mismatches; Go requires explicit error checking.
4. **No Undefined**: Go uses zero values instead of JavaScript's `undefined`.

---

## Next Steps

- Run `main.go` to see JSON decoding in action
- Modify the structs to include new fields and see how `Unmarshal()` handles them
- Try `json.Marshal()` to convert Go structs back to JSON
- Experiment with struct tags (`omitempty`, `-`, etc.)

---

## Further Reading

- [Go JSON Package Documentation](https://pkg.go.dev/encoding/json)
- [HTTP Spec (RFC 7231)](https://tools.ietf.org/html/rfc7231)
- [JSON Spec (RFC 8259)](https://tools.ietf.org/html/rfc8259)
