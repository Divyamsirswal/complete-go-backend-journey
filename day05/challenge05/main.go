package main

import "fmt"

/*
	Challenge 5 - Monk
	Function that:
		takes string
		returns reversed string
*/

func reverseString(str string) string {
	var str1 string
	for i := len(str) - 1; i >= 0; i-- {
		str1 += string(str[i])
	}
	return str1
}

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	fmt.Println("After Reversed:", reverseString(str))
}
