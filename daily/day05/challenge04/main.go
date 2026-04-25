package main

import "fmt"

/*
	Challenge 4: Factorial

	4! -> 1 x 2 x 3 x 4 => 24
	5! -> 120
*/

func findFactorial(num int) int {
	fact := 1

	for i := 1; i <= num; i++ {
		fact = fact * i
	}

	return fact
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	fmt.Println("Factorial of a number is:", findFactorial(num))
}
