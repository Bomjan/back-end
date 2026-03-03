// This file is the primary entry point for the 'mymath' module.
package main

import (
	"fmt"
	// Import the 'calculator' package from the 'mymath' module.
	// The import path is relative to the module root.
	"mymath/calculator"
)

func main() {
	fmt.Println("Package and Module Demonstration - Structural Organization")

	a, b := 10, 5

	// Call the 'Add' function from the imported 'calculator' package.
	// Note the capitalization: 'Add' is an exported (public) function.
	sum := calculator.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	// Call the 'Multiply' function from the 'calculator' package.
	product := calculator.Multiply(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, product)
}
