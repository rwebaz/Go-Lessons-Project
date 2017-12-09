package main

import (
	"fmt"
)

/* Original main function
func main() {
	fmt.Printf("Hello there, World!\n")
} */

// Note. You may declare and assign static variables regular way in Go ( var, name, type )

func main() {
	// Declare and assign variables
	var message string = (`Hello there, World!`)
	var newline string = ("\n")

	// Declare methods
	fmt.Printf(message + newline)
}
