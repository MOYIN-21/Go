package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter any letter and i will tell you if its alphabet or consonant")
	var letter string
	fmt.Scan(&letter)

	alphabet := character(letter)
	fmt.Println("letter" + " " + letter + " " + "is a" + " " + alphabet)
}

func character(letter string) string {
	if letter == "A" || letter == "E" || letter == "I" || letter == "0" || letter == "U" {
		return "Vowel"
	}
	return "Consonant"

}
