# Variables and Go's Type System

Master Go's strict typing, zero values, and the differences from JavaScript.

## Learning Objectives

- Understand Go's basic types: `int`, `float64`, `string`, `bool`
- Learn type inference with the `:=` operator
- Understand "zero values" (Go's alternative to `undefined`)
- Learn explicit type conversion
- Compare Go typing with JavaScript

## Go's Basic Types

Go is **statically typed**, meaning every variable has a fixed type determined at compile time.

### Numeric Types

**Integers** (whole numbers):
- `int` - The default integer type (size depends on your system)
- `int8`, `int16`, `int32`, `int64` - Fixed-size integers
- `uint` - Unsigned integer (no negative values)

**Floating Point**:
- `float32` - 32-bit decimal
- `float64` - 64-bit decimal (default for decimal literals like `3.14`)

```go
a := 42       // Type inferred as int
b := 3.14     // Type inferred as float64
c := int32(5) // Explicitly declare as int32
```

### Text and Boolean

**String**:
- `string` - Text enclosed in double quotes

**Boolean**:
- `bool` - `true` or `false`

```go
name := "Alice"      // string
isActive := true     // bool
```

## Variable Declaration

Go offers multiple ways to declare variables:

### 1. Short Declaration with Type Inference (`:=`)

```go
x := 10        // Type inferred as int
y := 3.14      // Type inferred as float64
z := "hello"   // Type inferred as string
```

**Only works inside functions!**

### 2. Explicit Declaration with `var`

```go
var x int = 10
var y float64 = 3.14
var z string = "hello"
```

### 3. Declaration with Type, No Value (Zero Value)

```go
var x int       // x = 0
var y float64   // y = 0.0
var z string    // z = ""
var b bool      // b = false
```

## Zero Values

Go has **no `undefined`**. Every type has a default "zero value":

| Type | Zero Value |
|------|------------|
| `int` | `0` |
| `float64` | `0.0` |
| `string` | `""` (empty string) |
| `bool` | `false` |

This eliminates a whole class of bugs that plague JavaScript!

```go
// JavaScript
let x;              // x is undefined
console.log(x + 5); // NaN (surprise!)

// Go
var x int           // x is 0
fmt.Println(x + 5)  // 5 (predictable!)
```

## Type Conversion

**Go does NOT auto-convert types.** You must be explicit.

### Correct: Explicit Conversion
```go
a := 10        // int
b := 3.14      // float64
sum := float64(a) + b  // Convert a to float64 first
// sum = 13.14
```

### Wrong: Implicit Conversion (Won't Compile)
```go
a := 10        // int
b := 3.14      // float64
sum := a + b   // ERROR: cannot add int and float64
```

## JavaScript vs Go: Typing

### JavaScript (Dynamic Typing)
```javascript
let x = 42;      // Number
x = "hello";     // String - OK, same variable
x = true;        // Boolean - OK, changes type freely

console.log(x + 5);   // "true5" - implicit conversion (confusing!)
```

### Go (Static Typing)
```go
var x int = 42
x = "hello"        // ERROR: cannot assign string to int

y := 42
y = y + 3.14       // ERROR: cannot add int and float64 (must convert)
```

## Printf Format Codes

When printing variables, use format specifiers:

```go
fmt.Printf("Integer: %d\n", 42)           // %d = integer
fmt.Printf("Float: %f\n", 3.14)           // %f = float
fmt.Printf("String: %s\n", "hello")       // %s = string
fmt.Printf("Boolean: %t\n", true)         // %t = boolean
fmt.Printf("Any type: %v\n", 42)          // %v = any value
fmt.Printf("Type info: %T\n", 42)         // %T = type of value
fmt.Printf("String with quotes: %q\n", "hi") // %q = quoted string
```

## Running This Example

```bash
go run main.go
```

You'll see:
1. Variable values and their inferred types
2. Zero values (defaults for uninitialized variables)
3. Explicit type conversion in action

## Key Takeaways

✅ Go is **statically typed** - type is fixed at compile time  
✅ Types are **explicit or inferred** - you choose clarity  
✅ **No auto-conversion** - prevents surprises  
✅ **Zero values instead of undefined** - safer code  
✅ **Explicit conversions** - `float64(intValue)`  

## JavaScript Developer Cheat Sheet

| JavaScript | Go | Notes |
|-----------|-----|--------|
| `let x = 5` | `x := 5` | Type inferred (int) |
| `let x = 5; x = "hi"` | `var x int = 5; x = "hi"` | ERROR - Go is strict |
| `console.log(5 + "5")` | `fmt.Println(5 + 5)` | No auto-conversion |
| `let x;` (undefined) | `var x int` (0) | Go has zero values |
| `5 + 3.14` | `float64(5) + 3.14` | Explicit conversion |

## Next Steps

1. Run the program and observe type inference
2. Try changing variable values and see what types are inferred
3. Modify the program to add more variables of different types
4. Try removing the explicit conversion and see the error
5. Move to the next module to learn about packages and imports

## Try It

Exercises:
```go
// 1. Declare an integer and float, then add them
a := 10
b := 5.5
c := float64(a) + b  // What's c?

// 2. Create string and number variables, observe zero values
var city string      // What's the zero value?
var population int   // What's the zero value?

// 3. Use %T to print the type of each variable
fmt.Printf("%T\n", a)
```

---

**Next Module**: [02 Intermediate - mymath_package](../../02_intermediate/mymath_package/)
