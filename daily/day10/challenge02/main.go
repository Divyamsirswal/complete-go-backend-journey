package main

import (
	"errors"
	"fmt"
)

/*
	Challenge 2: Even Number Check
	Function:
		- returns error if number is negative
*/

func checkEvenNumber(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Negative numbers are not allowed")
	}
	if n%2 == 0 {
		return 0, nil
	} else {
		return 1, nil
	}
}

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	res, err := checkEvenNumber(n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if res == 0 {
			fmt.Println("Number is even")
		} else {
			fmt.Println("Number is odd")
		}
	}

}
