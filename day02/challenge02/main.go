package main

import "fmt"

/*
	Challenge 2: Calculator
	Take:
		- two integers

	Print:
		- sum
		- difference
		- product

	Note: Use Printf
*/

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Printf("Addition : %d\nDifference : %d\nProduct : %d",
		(num1 + num2),
		(num1 - num2),
		(num1 * num2),
	)
}
