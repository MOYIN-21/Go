package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter score: ")
	var score int
	fmt.Scan(&score)

	grade := letterGrade(score)
	fmt.Print("The student grade is", grade)
}

func letterGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	default:
		return "F"

	}
}
