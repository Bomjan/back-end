package main

import "fmt"

func main() {
	// JS Connection: In JavaScript, we use 'let' or 'const' for any data.
	// In Go, we use specific types to help the computer manage memory efficiently.

	// 1. Numeric Types
	// 'int' is for whole numbers.
	// 'float64' is for decimals (JS equivalent to Number).
	a := 10        // inferred as int
	b := 5.5       // inferred as float64
	
	// 2. Text and Boolean
	// 'string' is for text.
	// 'bool' is for true/false.
	name := "Go Developer"
	isLearning := true

	fmt.Printf("Type: %T, Value: %v\n", a, a)
	fmt.Printf("Type: %T, Value: %v\n", b, b)
	fmt.Printf("Type: %T, Value: %v\n", name, name)
	fmt.Printf("Type: %T, Value: %v\n", isLearning, isLearning)

	// 3. Zero Values
	// JS Connection: In JS, an uninitialized variable is 'undefined'.
	// In Go, every type has a default "Zero Value".
	var defaultInt int       // 0
	var defaultFloat float64 // 0
	var defaultBool bool     // false
	var defaultString string // "" (empty string)

	fmt.Println("\nZero Values (Default Initializations):")
	fmt.Printf("int: %d, float: %f, bool: %t, string: %q\n", 
		defaultInt, defaultFloat, defaultBool, defaultString)

	// 4. Type Conversion
	// Unlike JS, Go will NOT automatically convert types.
	// You must explicitly convert if you want to add an int to a float.
	sum := float64(a) + b
	fmt.Printf("\nExplicit Conversion: %d converted to float + %f = %f\n", a, b, sum)
}
