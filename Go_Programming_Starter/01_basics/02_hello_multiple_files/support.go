// Part of 'package main', the main entry point package.
package main

import "fmt"

// Support is a secondary function used to demonstrate package-level scoping.
func support() {
	fmt.Println("Executing support function from support.go.")
}
