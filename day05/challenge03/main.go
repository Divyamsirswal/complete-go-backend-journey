package main

import "fmt"

/*
	Challenge 3: Max of Two Numbers
	return the max of 2 int value and return it
*/

func findMaximum(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Maximum number is:", findMaximum(a, b))
}
