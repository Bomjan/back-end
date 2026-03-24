// This file demonstrates Go Modules and package imports.
package main

import (
	"fmt"
	"mymath/calculator"
)

func main() {
	fmt.Println("=== Package and Module Demonstration ===\n")

	// Example 1: Basic arithmetic with exported functions
	fmt.Println("Example 1: Calling exported functions from 'calculator' package")
	a, b := 10, 5
	sum := calculator.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	product := calculator.Multiply(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, product)

	// Example 2: Working with slices
	fmt.Println("\nExample 2: Creating and manipulating slices")
	sliceName := make([]int, 3, 5)
	sliceName = append(sliceName, 34, 5, 6, 9)
	sliceName[0] = 9
	fmt.Println("Modified slice:", sliceName)

	// Example 3: Array slicing
	fmt.Println("\nExample 3: Array and slice operations")
	myarr := [10]int{4, 6, 7, 1, 4, 34, 67, 45, 23, 88}
	newSlice := myarr[3:9]
	fmt.Println("Full array: ", myarr)
	fmt.Println("Slice [3:9]:", newSlice)

	// Example 4: Slice literals
	fmt.Println("\nExample 4: Slice operations")
	mySlice := []int{5, 3, 6, 7}
	fmt.Println("Initial slice:", mySlice)

	s := []int{1, 7, 5, 9, 2, 4, 0}
	s = append(s, 67)
	fmt.Println("Slice after append:", s)
}
