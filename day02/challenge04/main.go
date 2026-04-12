package main

import "fmt"

/*
	Challenge 4 - Monk

	Take:
		- marks

	Output:
		- = 75 → Distinction
		- 60–74 → First Class
		- 40–59 → Pass
		- < 40 → Fail
*/

func main() {
	var marks int

	fmt.Print("Enter your marks: ")
	fmt.Scan(&marks)

	if marks >= 75 {
		fmt.Println("Distinction")
	} else if marks >= 60 && marks <= 74 {
		fmt.Println("First Class")
	} else if marks >= 40 && marks <= 59 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
