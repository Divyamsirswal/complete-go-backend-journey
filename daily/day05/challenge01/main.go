package main

import "fmt"

/*
	Challenge 1: Simple Calculator Functions
	Create functions:
		add(a, b)
		subtract(a, b)
		multiply(a, b)
	Call them in  -> main()
*/

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Addition:", add(a, b))
	fmt.Println("Subtraction:", subtract(a, b))
	fmt.Println("Multiplication:", multiply(a, b))
}
