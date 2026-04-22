package main

import (
	"errors"
	"fmt"
)

/*
	Challenge 3: User Input Validation
	Take input:
		- age
			- if age < 0 → error
			- else print valid
*/

func checkUserAge(age int) (string, error) {
	if age < 0 {
		return "invalid", errors.New("Age cannot be negative")
	}

	return "valid", nil
}

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	res, err := checkUserAge(age)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}
