package main

import "fmt"

/*
	Challenge 4: Multiplication Table
	Take input:
		- n (number)
	Output:
		n x 1 = x
		n x 2 = x
		...
		n x 10 = x
*/

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Println(n, " x ", i, " = ", i*n)
	}
}
