// Part of 'package main', which is the executable package for this application.
package main

import "fmt"

// Helper is a function accessible to other files within the same package.
func helper() {
	fmt.Println("Executing helper function from helper.go.")

	// Call the support function defined in support.go.
	support()
}
