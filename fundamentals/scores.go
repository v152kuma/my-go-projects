package main

import (
	"fmt"
	"strconv"
)

func commandLineScoreGetter() {

	type score struct {
		name  string
		score int
	}

	scores := []score{} // Initialize an empty slice of score structs

	for {
		// Read input from the user
		var name string
		var rawScore string
		fmt.Print("Enter student name (or 'exit' to quit): ")
		fmt.Scanln(&name)
		if name == "exit" {
			break
		}
		fmt.Print("Enter student score: ")
		fmt.Scanln(&rawScore)
		s, _ := strconv.Atoi(rawScore) // Convert string to int

		// Append the new score to the slice
		scores = append(scores, score{name, s})
	}

	for _, student := range scores {
		fmt.Println(student.name, student.score)
	}

}
