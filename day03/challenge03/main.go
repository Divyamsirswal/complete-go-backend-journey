package main

import "fmt"

/*
	Challenge 3: Sum of Numbers
	Take input:
		- n (int)
	Print:
		- sum of 1 to n
*/

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("Sum : ", sum)
}
