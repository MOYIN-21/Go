package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Enter any number")
	fmt.Scan(&x)
	fmt.Print("Enter any number")
	fmt.Scan(&y)
	if x == y {
		fmt.Println("0")
	} else if x > y {
		fmt.Println("1")
	} else if x < y {
		fmt.Println("-1")
	} else {
		fmt.Println("invalid input")
	}

}