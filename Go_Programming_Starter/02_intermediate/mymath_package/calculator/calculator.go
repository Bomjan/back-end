// Package calculator provides fundamental arithmetic operations.
// This package is part of the 'mymath' module.
package calculator

// Add is an exported function that performs integer addition.
// Exporting is accomplished by capitalizing the first letter of the identifier.
func Add(a int, b int) int {
	return a + b
}

// Multiply is an exported function that performs integer multiplication.
// It is accessible to other packages that import 'package calculator'.
func Multiply(a int, b int) int {
	return a * b
}

// Unexported identifiers (starting with a lowercase letter) are private to this package.
func subtract(a int, b int) int {
	return a - b
}
