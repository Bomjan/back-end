// This file is part of 'package main'.
// Multiple files in the same directory can share the same package name and access each other's identifiers.
package main

import "fmt"

func main() {
	fmt.Println("Program execution started in main.go.")

	// Call the helper function defined in helper.go.
	// Since both files belong to the same package, no explicit import is required.
	helper()
}
