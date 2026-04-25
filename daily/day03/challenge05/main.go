package main

import "fmt"

/*
	Challenge 5 - Monk
	Print this pattern:
		*
		**
		***
		****
		*****
*/

func main() {

	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
