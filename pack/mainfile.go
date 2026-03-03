package main

import (
	"fmt"
	"pack_demo/calculator" // instead of pack_demo write mymath
)

func main() {
	fmt.Println("Package and Module Demonstration")

	a, b := 10, 5

	sum := calculator.Add(a, b) // instead of Add write add
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	product := calculator.Multiply(a, b) // similar to line 13
	fmt.Printf("%d * %d = %d\n", a, b, product)
}
