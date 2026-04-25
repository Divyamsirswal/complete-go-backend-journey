package main

import "fmt"

/*
	Challenge 3: Pattern

	*****
	****
	***
	**
	*
*/

func main() {

	for i := 1; i <= 5; i++ {
		for j := i; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
