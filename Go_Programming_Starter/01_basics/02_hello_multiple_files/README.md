# Hello, Multiple Files!

Learn how Go organizes code using **packages** and file structure.

## Learning Objectives

- Understand package-level scope in Go
- Organize code across multiple files
- Distinguish between exported (public) and unexported (private) identifiers
- Learn Go's visibility rules based on capitalization

## Key Concepts

### Package Scope

In Go, all files in the **same directory** belong to the **same package**. They share scope automatically—no `export` keyword needed!

```
02_hello_multiple_files/
├── main.go          # package main
├── helper.go        # also package main
└── support.go       # also package main
```

All three files can call functions from each other without any imports.

### Exported vs Unexported

Go uses **capitalization** to control visibility:

- **Capitalized** identifiers (`HelloWorld`, `Add`, `Person`) are **exported** (public)
- **lowercase** identifiers (`helloWorld`, `add`, `person`) are **unexported** (private)

```go
// In helper.go
func ExportedFunction() {
    // Other packages can call this
}

func unexportedFunction() {
    // Only this package can call this
}
```

This is different from JavaScript:

| Concept | JavaScript | Go |
|---------|-----------|-----|
| Export syntax | `export default`, `export const` | Capitalization (automatic) |
| Private syntax | No keyword (just don't export) | lowercase (automatic) |
| Scope | Module-level | Package-level (same directory) |

### Example: Three Files in One Package

**main.go**:
```go
package main

func main() {
    greeting := getGreeting()  // Calls unexported function
    PrintGreeting(greeting)    // Calls exported function from helper.go
}
```

**helper.go**:
```go
package main

func getGreeting() string {  // lowercase = unexported (private to package)
    return "Hello from helper!"
}

func PrintGreeting(msg string) {  // Capitalized = exported (but still package-scoped)
    fmt.Println(msg)
}
```

## This Example

This directory contains three files:
- **main.go** - Entry point with main()
- **helper.go** - Provides helper functions
- **support.go** - Provides support functions

All three are in the `main` package and can call each other's functions.

## Running This Example

```bash
go run .
# Runs all .go files in the directory

# Or explicitly:
go run main.go helper.go support.go
```

## Key Differences from JavaScript

### JavaScript Module System
```javascript
// file1.js
export function add(a, b) { return a + b; }

// file2.js
import { add } from './file1.js';
console.log(add(2, 3));  // Must explicitly import
```

### Go Package System
```go
// file1.go
package main
func Add(a int, b int) int { return a + b }

// file2.go
package main
func main() {
    Add(2, 3)  // No import needed! Same package = same scope
}
```

## Important Notes

1. **All files in a directory must be in the same package** - You cannot have `package main` in one file and `package utils` in another file in the same directory

2. **Capitalization is the only visibility mechanism** - There's no `private` keyword; Go relies on naming conventions

3. **Package vs Module** - A package is a directory of `.go` files; a module is a collection of packages (defined in `go.mod`)

## Next Steps

1. Run this program to see how multiple files work together
2. Add your own function in `support.go` (capitalized to make it exported)
3. Call it from `main.go`
4. Move to `03_variables/` to learn about Go's type system

## Try It

Exercises:
1. Add a new exported function in `helper.go` and call it from `main()`
2. Create a function that uses both `helper.go` and `support.go` functions
3. Try creating a lowercase function and see if you can call it from a different file (you should be able to within the same package)

---

**Next Module**: [03 Variables](../03_variables/)
