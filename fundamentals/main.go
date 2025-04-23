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

	//name, score := "John Doe", 95.5 // Short variable declaration

	students := []string{"John Doe", "Jane Smith", "Alice Johnson"}
	scores := []float64{95.5, 88.0, 92.5}
	// Using a for loop to iterate over the slices
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	for i := 0; i < len(students); i++ {
		name := students[i]
		score := scores[i]
		// Print a line of dashes from strings package
		fmt.Println(name, score)
	}

	studentScoresMap := map[string]float64{
		"John Doe":      95.5,
		"Jane Smith":    88.0,
		"Alice Johnson": 92.5,
	}
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	for name, score := range studentScoresMap {
		fmt.Println(name, score)
	}

	type score struct {
		name  string
		score float64
	}

	studentScores := []score{
		{"John Doe", 95.5},
		{"Jane Smith", 88.0},
		{"Alice Johnson", 92.5},
	}

	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14))
	for _, student := range studentScores {
		fmt.Println(student.name, student.score)
	}
	// Using a for loop to iterate over the slices

}
