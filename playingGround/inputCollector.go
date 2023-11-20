package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	if input%2 == 0 {
		fmt.Println(input, "even")
	} else {
		fmt.Println(input, "odd")
	}
}