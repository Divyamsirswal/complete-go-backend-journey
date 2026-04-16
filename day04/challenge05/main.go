package main

import "fmt"

/*
	Challenge 5: Check Palindrome String or not
	Input:  madam
	Output: Palindrome
*/

func main() {

	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	var i, j int
	i = 0
	j = len(str) - 1

	for i <= j {
		if rune(str[i]) != rune(str[j]) {
			fmt.Println("Not Palindrome")
			return
		}
		i++
		j--
	}

	fmt.Println("Palindrome")
}
