// This file is the primary entry point for the 'mymath' module.
package main

import "fmt"

func main() {
	// fmt.Println("Package and Module Demonstration - Structural Organization")

	// a, b := 10, 5

	// Call the 'Add' function from the imported 'calculator' package.
	// Note the capitalization: 'Add' is an exported (public) function.
	// sum := calculator.Add(a, b)
	// fmt.Printf("%d + %d = %d\n", a, b, sum)

	// Call the 'Multiply' function from the 'calculator' package.
	// product := calculator.Multiply(a, b)
	// fmt.Printf("%d * %d = %d\n", a, b, product)

	// sliceName := make([]int, 3, 5)
	// sliceName = append(sliceName, 34, 5, 6, 9)
	// sliceName[0] = 9

	// fmt.Println(sliceName)

	// myarr := [10]int{4, 6, 7, 1, 4, 34, 67, 45, 23, 88}
	// newSlice := myarr[3:9]
	// fmt.Println(myarr)
	// fmt.Println((newSlice))

	// mySlice := []int{5, 3, 6, 7}
	// fmt.Println(mySlice)

	// s := []int{1, 7, 5, 9, 2, 4, 0}
	// s = append(s, 67)
	// fmt.Println(s)

	s := make([]int, 3, 5)

	fmt.Println(s)

}
