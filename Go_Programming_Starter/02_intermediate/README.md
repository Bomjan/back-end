# Module 02: Intermediate Go Concepts

Welcome to the intermediate section of your Go learning journey. Here, you'll move beyond the basics and into more sophisticated Go patterns, package management, and real-world applications.

## 📚 Curriculum Overview

This module is organized into four key learning areas:

### 1. [mymath_package](./mymath_package/) - Go Modules and Custom Packages
Learn how to organize Go code into reusable, importable packages.

**Topics**:
- Go modules and `go.mod`
- Creating custom packages
- Identifier visibility (exported vs unexported)
- Module-based project structure
- The "Capitalization Rule"

**Example**: A calculator package demonstrating imports, slices, and arrays

**Key Files**:
- `mymath_package/go.mod` - Module declaration
- `mymath_package/calculator/calculator.go` - Custom package with exported functions
- `mymath_package/main.go` - Importing and using the calculator package

---

### 2. [HTTP](./HTTP/) - JSON Encoding, Decoding & API Responses
Master Go's approach to JSON serialization and data marshaling.

**Topics**:
- HTTP fundamentals (requests, responses, methods)
- JSON encoding with `json.Marshal()`
- JSON decoding with `json.Unmarshal()`
- Struct tags (`json:"..."`, `omitempty`, `-`)
- Nested structures and arrays
- Error handling in JSON operations
- Real-world API response patterns

**Example**: Seven comprehensive examples showing JSON operations, struct tags, pretty printing, error handling, and API responses

**Key Concepts**:
- Encoding: Go struct → JSON bytes (for sending data)
- Decoding: JSON bytes → Go struct (for receiving data)
- Marshaling: `json.Marshal()`, `json.MarshalIndent()`
- Unmarshaling: `json.Unmarshal()`

---

### 3. [presentation](./presentation/) - Pointers, Maps & Real-World Applications
Learn about Go's pointer system and practical data structures.

**Topics**:
- Pointers and memory addresses (`&` and `*` operators)
- Maps (key-value storage)
- Slices of structs
- Removing items from slices
- Using pointers in functions
- Real-world API patterns with structs

**Example**: Demonstrates pointer usage, maps, and real-world User struct patterns

**Key Insight**: Understand how Go manages memory references vs JavaScript's implicit object references

---

## 🎯 Learning Path

### Recommended Order:
1. **Start here**: [`mymath_package`](./mymath_package/) - understand module structure
2. **Then**: [`HTTP`](./HTTP/) - work with real data (JSON)
3. **Finally**: [`presentation`](./presentation/) - advanced patterns (pointers, maps)

### Running Examples

Each subdirectory contains a complete, runnable example:

```bash
# Run mymath_package
cd mymath_package && go run .

# Run HTTP examples
cd HTTP && go run main.go

# Run presentation examples
cd presentation && go run main.go
```

---

## 🔑 Key Concepts Across This Module

### Packages & Modules
- **Package**: A directory of Go files
- **Module**: A collection of packages managed by `go.mod`
- **Visibility**: Capitalization determines if identifiers are exported

### JSON & Data Serialization
- **Struct tags**: Control JSON field names and behavior
- **Marshal**: Convert Go → JSON bytes
- **Unmarshal**: Convert JSON bytes → Go

### Pointers & References
- **`&`**: Get the memory address of a variable
- **`*`**: Dereference to access the value at an address
- **Efficiency**: Use pointers to avoid copying large structs

### Maps & Slices
- **Maps**: Key-value storage with typed keys and values
- **Slices**: Dynamic arrays with append/slice operations

---

## 📊 Quick Comparison: JavaScript vs Go

### Modules & Organization

**JavaScript**:
```javascript
// file1.js
export function add(a, b) { return a + b; }

// file2.js
import { add } from './file1.js';
```

**Go**:
```go
// calculator/calculator.go
package calculator
func Add(a, b int) int { return a + b }

// main.go
package main
import "myapp/calculator"
func main() {
    calculator.Add(5, 3)
}
```

### Working with JSON

**JavaScript**:
```javascript
const obj = { name: "Alice", age: 30 };
const json = JSON.stringify(obj);          // Object → JSON string
const parsed = JSON.parse(json);           // JSON string → Object
```

**Go**:
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
obj := Person{Name: "Alice", Age: 30}
json, _ := json.Marshal(obj)               // Struct → JSON bytes
json.Unmarshal(json, &obj)                 // JSON bytes → Struct
```

### References & Pointers

**JavaScript**:
```javascript
let a = { value: 42 };
let b = a;              // b references same object
b.value = 100;
console.log(a.value);   // 100 (both reference same object)
```

**Go**:
```go
type Data struct { Value int }
a := Data{Value: 42}
b := a                  // b is a COPY
b.Value = 100
fmt.Println(a.Value)    // 42 (unchanged)

// To reference same data:
ptr := &a
ptr.Value = 100
fmt.Println(a.Value)    // 100 (changed through pointer)
```

---

## 📝 Best Practices Covered

✅ **Organize code into packages** for reusability  
✅ **Use struct tags** to control JSON serialization  
✅ **Handle errors explicitly** (especially with Marshal/Unmarshal)  
✅ **Use pointers efficiently** to avoid copying large data  
✅ **Follow Go naming conventions** (PascalCase for exported, camelCase for unexported)  
✅ **Type your maps** (specify key and value types)  

---

## 🚀 What You'll Be Able to Do

After completing this module, you'll be able to:

- ✅ Create and import custom packages
- ✅ Convert between Go structs and JSON
- ✅ Build API response structures
- ✅ Use maps and pointers effectively
- ✅ Organize projects using Go modules
- ✅ Handle real-world data serialization
- ✅ Apply Go's visibility rules correctly

---

## 📂 Directory Structure

```
02_intermediate/
├── README.md                    # This file
├── mymath_package/              # Modules & packages
│   ├── go.mod
│   ├── main.go
│   ├── README.md
│   └── calculator/
│       └── calculator.go
├── HTTP/                        # JSON encoding/decoding
│   ├── main.go
│   └── README.md
└── presentation/                # Pointers & maps
    ├── main.go
    └── README.md
```

---

## 🎓 Next Steps

1. **Master Modules**: Complete `mymath_package/` and understand how to structure Go projects
2. **Work with Data**: Complete `HTTP/` and learn JSON serialization
3. **Advanced Patterns**: Complete `presentation/` and learn pointers and maps
4. **Build Something**: Combine these concepts to build a small project (e.g., a simple API server)

---

## 📖 Further Reading

- [Go Packages](https://golang.org/doc/code#PackageNames)
- [JSON Package Documentation](https://pkg.go.dev/encoding/json)
- [Pointers in Go](https://golang.org/doc/effective_go#pointers_vs_values)
- [Go Modules](https://golang.org/doc/modules/gomod-ref)

---

**Previous Module**: [01 Basics](../01_basics/)  
**Next Module**: Coming soon - Error Handling, Interfaces, Goroutines
