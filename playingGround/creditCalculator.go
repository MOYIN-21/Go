package main

import "fmt"

func main() {
	var input int
	var beginning int
	var items int
	var total int
	var credit int
	fmt.Println("Enter your account number: ")
	fmt.Scan(&input)
	fmt.Println("balance at the beginning of the month: ")
	fmt.Scan(&beginning)
	fmt.Println("total of all items charged by the customer this month:")
	fmt.Scan(&items)
	fmt.Println("total of all credits applied to the customer’s account this month: ")
	fmt.Scan(&total)
	fmt.Println("allowed credit limit.")
	fmt.Scan(&credit)

}
