package main

import "fmt"

/*
	Challenge 2: Even or Odd Function
	Return:
		true → even
		false → odd
*/

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	res := isEven(num)
	if res == true {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}
}
