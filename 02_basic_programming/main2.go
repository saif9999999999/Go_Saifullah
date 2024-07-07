package main

import (
	"fmt"
)

func main() {
	score := 78 // Replace with any score you want to test

	// Define the grade based on the score range
	var grade string
	switch {
	case score >= 85 && score <= 100:
		grade = "A"
	case score >= 70 && score <= 84:
		grade = "B"
	case score >= 55 && score <= 69:
		grade = "C"
	case score >= 40 && score <= 54:
		grade = "D"
	case score >= 0 && score <= 39:
		grade = "E"
	default:
		grade = "Invalid score"
	}

	// Print the result
	fmt.Printf("Score: %d, Grade: %s\n", score, grade)
}
