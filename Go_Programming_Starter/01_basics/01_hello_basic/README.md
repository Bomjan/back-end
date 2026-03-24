# Hello, Go World!

Your first Go program—the entry point to learning Go.

## Learning Objectives

- Understand the `main` package and `main()` function
- Learn how to import and use the `fmt` package
- Write your first executable Go program
- Run Go code with `go run`

## Key Concepts

### The `main` Package

Every executable Go program must have:
1. A package called `main`
2. A function called `func main()`

When you run `go run main.go`, Go looks for the `main()` function and executes it.

```go
package main  // This declares a package

func main() {  // This is the entry point
    // Your code here
}
```

### The `fmt` Package

`fmt` is the **formatting** package in Go's standard library. It provides functions to print output:

- `fmt.Println()` - prints text followed by a newline
- `fmt.Print()` - prints text without a newline  
- `fmt.Printf()` - prints formatted text (like printf in C or JavaScript's template strings)

```go
fmt.Println("Hello, World!")     // Prints: Hello, World!
fmt.Print("No newline")          // Prints: No newline (no \n at end)
fmt.Printf("Number: %d\n", 42)   // Prints: Number: 42
```

## Running This Example

```bash
go run main.go
# Output: Hello, Go World!
```

## What's Different from JavaScript?

| Aspect | JavaScript | Go |
|--------|-----------|-----|
| Output | `console.log()` | `fmt.Println()` |
| Package system | `require()` or `import` statements | `package` + `import` blocks |
| Entry point | Global scope or `main()` function | Must have `func main()` |
| Execution | Node.js runtime | `go run` or compiled binary |
| Type system | Dynamic (inferred at runtime) | Static (declared at compile time) |

## Next Steps

1. Modify the program to print different messages
2. Use `fmt.Printf()` to print formatted output
3. Move to `02_hello_multiple_files/` to learn package scope and organizing code across files

## Try It

Try these exercises:

```go
// Print your name
fmt.Println("My name is ...")

// Print a calculation
fmt.Printf("2 + 2 = %d\n", 2+2)

// Print multiple values on one line
fmt.Println("Hello", "World", "!")
```

---

**Next Module**: [02 Hello Multiple Files](../02_hello_multiple_files/)
