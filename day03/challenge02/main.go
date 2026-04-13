package main

import "fmt"

/*
	Challenge 2: Even Numbers
	Print even numbers from 1 to 10
*/

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
