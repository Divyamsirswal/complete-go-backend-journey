package main

import "fmt"

/*
	Challenge 3: Pass/Fail System

	Take:
		- marks (int)

	Logic:
		- = 40 → "Pass"
		- < 40 → "Fail"
*/

func main() {
	var marks int

	fmt.Print("Enter your marks: ")
	fmt.Scan(&marks)

	if marks < 40 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Pass")
	}
}
