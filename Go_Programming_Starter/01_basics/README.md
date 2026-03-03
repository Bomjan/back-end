# Module 01: Foundational Concepts

Welcome to the foundations of Go. This section covers the core syntax and the initial steps to building a Go application.

## Learning Objectives

- **[01_hello_basic](./01_hello_basic)**: Introduction to the `main` package, the entry point of a Go application, and the standard `fmt` library.
- **[02_hello_multiple_files](./02_hello_multiple_files)**: Implementation of package-level scope and the organization of multiple files within a single package.
- **[03_variables](./03_variables)**: Statically typed variable declaration, initialization, and usage.

---

## Technical Overview

### 1. The `main` Package
In Go, every executable program must be part of `package main`. This informs the Go compiler that this codebase is an application rather than a library. The execution always begins at the `func main()`.

### 2. The `fmt` Standard Library
The `fmt` package provides standard input/output formatting. The most common tool is `Println`, which outputs formatted text followed by a newline.

### 3. Package-Level Scope
Files within the same directory must share the same package name. Identifiers (functions, variables, etc.) defined in one file are visible to all other files within that same package without any explicit imports.

---

## JavaScript to Go: Connecting the Dots

If you are coming from a JavaScript background, Go's type system might feel strict. Here is how to map your JS knowledge to Go:

### Type Mapping Table

| Feature | JavaScript (ES6+) | Go (Golang) | Key Difference |
| :--- | :--- | :--- | :--- |
| **Declaration** | `let`, `const`, `var` | `var`, `:=` | Go is statically typed; types are checked at compile time. |
| **Numbers** | `Number`, `BigInt` | `int`, `float64`, `int64`, etc. | Go has specific sizes for numbers to optimize memory. |
| **Decimals** | `Number` (64-bit float) | `float64`, `float32` | No single "Number" type; you must be explicit. |
| **Empty Value** | `undefined`, `null` | **Zero Values** (`0`, `""`, `false`) | Go variables always have a predictable default value. |
| **Strings** | `"`, `'`, \`` | `"`, \`` | Use `"` for standard strings and \`` for multi-line/raw strings. |

### The Numeric System
In JavaScript, all numbers are technically 64-bit floats. This means `10` and `10.5` are the same internal type.

In Go, **Integers (`int`)** and **Floats (`float64`)** are distinct.
- You cannot add an `int` to a `float64` without an explicit conversion: `float64(myInt) + myFloat`.
- This prevents precision bugs that are common in dynamically typed languages.

### Zero Values vs Undefined
In JavaScript, if you declare `let x;`, its value is `undefined`.
In Go, if you declare `var x int`, its value is `0`. 

There is no "undefined" in Go. This makes your code safer by ensuring variables always start in a known state.

---

## Execution Instructions

1.  **Navigate** to the specific concept directory using your terminal.
2.  **Execute** the program using the Go command-line tool:
    ```bash
    # For a single-file application:
    go run main.go
    
    # For multiple files within the same package:
    go run *.go
    ```
