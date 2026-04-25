package main

import "fmt"

/*
	Challenge 4: Number Pattern
	1
	12
	123
	1234
	12345
*/

func main() {

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
