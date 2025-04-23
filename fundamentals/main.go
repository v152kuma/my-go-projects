// main means entry point of the program
package main

import (
	"fmt"
	"strings"
)

func main() {

	const (
		// Constants are immutable values
		PI = 3.14
		// Untyped constants can be assigned to any type
		b string = "Hello, World!"
		// Typed constants must be assigned to a specific type
		c int32 = 42
	)

	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(PI)

	name, score := "John Doe", 95.5 // Short variable declaration
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14)) // Print a line of dashes from strings package
	fmt.Println(name, score)

}
