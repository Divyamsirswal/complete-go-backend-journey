package main

import (
	"errors"
	"fmt"
)

/*
	Challenge 4: String Length Check
	Function:
		- takes string
		- returns error if empty
*/

func checkStringLength(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("string can't be empty")
	}
	return str, nil
}

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&str)

	res, err := checkStringLength(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res, "is of", len(res), "size.")
	}
}
